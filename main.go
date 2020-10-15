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
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Title    string //标题
	From     string //发件人
	Subject  string //邮件标题
	Text     string //邮件正文
	HTML     string //邮件正文 html格式
	SMTP     string //smtp地址
	Port     int    //端口
	SMTPPort int    //发送端口
	Username string //邮件用户名
	Password string //邮件密码
}

func main() {
	//监听resume.md的变化
	fileWatcher.WatchResume(initProject)
}

func initProject() {
	//加载配置文件
	conf := loadConfig()
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

func initIndex(router *gin.Engine, conf Config) {
	//加载html模板
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)
	router.GET("/", func(c *gin.Context) {
		tmpl := loadStringFile("./assets/resume.tmpl")
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"title":  conf.Title,
			"resume": template.HTML(tmpl),
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

func initRouter(router *gin.Engine, e *gomail.Dialer, conf Config) {
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

func loadConfig() Config {
	filePtr, _ := os.Open("./config.json")
	defer filePtr.Close()
	var conf Config
	decoder := json.NewDecoder(filePtr)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("加载配置文件失败", err.Error())
	} else {
		fmt.Println("加载配置文件成功")
	}
	return conf
}

func initEmail(conf Config) *gomail.Dialer {
	e := gomail.NewDialer(conf.SMTP, conf.SMTPPort, conf.Username, conf.Password)
	//使用ssl
	e.SSL = true
	return e
}

func sendEmail(conf Config, address string, e *gomail.Dialer) error {
	//发送邮件
	mail := gomail.NewMessage()
	mail.SetHeader("From", conf.Username)
	mail.SetHeader("To", address)
	mail.SetHeader("Subject", conf.Subject)
	mail.SetBody("text/plain", conf.Text)
	mail.SetBody("text/html", conf.HTML)
	mail.Attach("./assets/resume.pdf", gomail.Rename(conf.Title+".pdf"))
	err := e.DialAndSend(mail)
	return err
}

func loadStringFile(p string) string {
	file, _ := os.Open(p)
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	return string(content)
}
