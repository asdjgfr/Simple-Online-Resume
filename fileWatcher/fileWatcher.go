package fileWatcher

import (
	"bufio"
	"bytes"
	"github.com/bep/debounce"
	"github.com/fsnotify/fsnotify"
	"io"
	"log"
	"myModule/types"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Cb func(types.Config)

func WatchResume(cb Cb, conf types.Config) {
	debounced := debounce.New(100 * time.Millisecond)
	//项目启动时重新生成resume
	debounced(func() {
		//100毫秒内只触发一次
		rebuildResume(conf)
	})
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("新建错误：", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					debounced(func() {
						rebuildResume(conf)
					})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("监听错误：", err)
			}
		}
	}()
	err = watcher.Add("./assets/resume.md")
	if err != nil {
		log.Fatal("添加监听错误：", err)
	}
	//回调事件
	cb(conf)
	<-done
}

func rebuildResume(conf types.Config) {
	//替换路径
	var resumePath="./assets/resume.md"
	err:= replacePath(resumePath)
	if err!=nil {
		log.Fatalf("替换路径失败： %s\n", err)
	}
	//重新生成html和pdf
	cmd := exec.Command("bash", "-c", "./bin/pandoc "+resumePath+".tmp.md -t html -o ./assets/resume.tmpl ; ./bin/pandoc --highlight-style zenburn "+resumePath+".tmp.md --pdf-engine=xelatex -t latex -V CJKmainfont='"+conf.Font+"' -V colorlinks -V urlcolor=NavyBlue -o ./assets/resume.pdf")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("生成resume失败： %s\n", err)
	}
	log.Println("生成resume成功")
	_ = os.Remove(resumePath+".tmp.md")
}

func replacePath(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	var result bytes.Buffer
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		result.WriteString(strings.Replace(string(a), "./", "./assets/", -1 ))
		result.WriteString("\n")
	}

	fw, err := os.OpenFile(path+".tmp.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	w := bufio.NewWriter(fw)
	_, err =w.WriteString(result.String())
	if err != nil {
		return err
	}
	err=w.Flush()
	return err
}
