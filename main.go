package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"html/template"
	"io/ioutil"
	"myModule/bindata"
	"myModule/fileWatcher"
	"myModule/types"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//加载配置文件
	Conf := loadConfig()
	//监听input.md的变化
	fileWatcher.WatchResume(initProject, Conf)
}

func initProject(conf types.Config) {
	//初始化email
	e := initEmail(conf)
	//初始化gin
	router := gin.Default()
	initIndex(router, conf)
	//加载静态文件
	router.Static("/assets", "./assets")
	//初始化路由
	initRouter(router, e, conf)
	err := router.Run(":" + strconv.Itoa(conf.Port))
	if err != nil {
		fmt.Println("启动失败：", err)
	}
}

func initIndex(router *gin.Engine, conf types.Config) {
	//加载html模板
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	router.GET("/", func(c *gin.Context) {
		tmpl := loadStringFile("./assets/input.tmpl")
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"title":  conf.Title,
			"content": template.HTML(tmpl),
			"i18n":   conf.I18n[conf.Language],
			"scripts":conf.Scripts,
		})
	})
}

// loadTemplate 加载由 go-assets-builder 嵌入的模板
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range bindata.Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func initRouter(router *gin.Engine, e *gomail.Dialer, conf types.Config) {
	addressMap := map[string]time.Time{}
	ipMap := map[string]int{}
	//发送邮件
	router.POST("/api/send-email", func(context *gin.Context) {
		clientIP := context.ClientIP()
		address := context.PostForm("address")
		if _, hasAddress := addressMap[address]; hasAddress && time.Now().Sub(addressMap[address]).Minutes() < 1 {
			//存在历史邮件地址
			context.JSON(http.StatusOK, gin.H{
				"msg":    "发送间隔小于1分钟，请稍后再试！",
				"status": 0,
			})
			return
		}
		if _, hasIP := ipMap[clientIP]; hasIP && ipMap[clientIP] > 9 {
			//ip发送次数
			context.JSON(http.StatusOK, gin.H{
				"msg":    "同一IP一天最多发送10次！",
				"status": 0,
			})
			return
		}
		if address == "" {
			context.JSON(http.StatusMethodNotAllowed, gin.H{
				"msg":    "未获取到地址",
				"status": 0,
			})
			return
		}
		err := sendEmail(conf, address, e)
		addressMap[address] = time.Now()
		ipMap[context.ClientIP()] = ipMap[context.ClientIP()] + 1
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg":    "发送失败！" + err.Error(),
				"status": 0,
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg":    "发送成功！",
				"status": 1,
			})
		}
	})
}

func loadConfig() types.Config {
	filePtr, _ := os.Open("./config.json")
	defer filePtr.Close()
	var conf types.Config
	decoder := json.NewDecoder(filePtr)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("加载配置文件失败", err.Error())
	} else {
		fmt.Println("加载配置文件成功")
	}
	return conf
}

func initEmail(conf types.Config) *gomail.Dialer {
	e := gomail.NewDialer(conf.SMTP, conf.SMTPPort, conf.Username, conf.Password)
	//使用ssl
	e.SSL = true
	return e
}

func sendEmail(conf types.Config, address string, e *gomail.Dialer) error {
	//发送邮件
	mail := gomail.NewMessage()
	mail.SetHeader("From", conf.Username)
	mail.SetHeader("To", address)
	mail.SetHeader("Subject", conf.Subject)
	mail.SetBody("text/plain", conf.Text)
	mail.SetBody("text/html", conf.HTML)
	mail.Attach("./assets/"+conf.Title+".pdf", gomail.Rename(conf.Title+".pdf"))
	err := e.DialAndSend(mail)
	return err
}

func loadStringFile(p string) string {
	file, _ := os.Open(p)
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	return string(content)
}
