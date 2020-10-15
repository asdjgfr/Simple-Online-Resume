package fileWatcher

import (
	"github.com/bep/debounce"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"time"
)

type Cb func()

func WatchResume(cb Cb) {
	debounced := debounce.New(100 * time.Millisecond)
	//项目启动时重新生成resume
	debounced(rebuildResume)
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
					//100毫秒内只触发一次
					debounced(rebuildResume)
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
	cb()
	<-done
}

func rebuildResume()  {
	//重新生成html和pdf
	cmd := exec.Command("./bin/pandoc", "./assets/resume.md","-o","assets/resume.tmpl")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
