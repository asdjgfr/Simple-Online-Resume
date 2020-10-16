package bindata

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\r\n<html lang=\"zh-CN\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n    <meta name=\"apple-mobile-web-app-capable\" content=\"yes\"/>\r\n    <link rel=\"shortcut icon\" href=\"/assets/favicon.ico\">\r\n    <title>{{.title}}</title>\r\n    <link rel=\"stylesheet\" href=\"/assets/style.css\">\r\n</head>\r\n<body>\r\n<header>\r\n    <div class=\"header\" id=\"header\">\r\n        <strong>如果遇到排版混乱，下载失败，发送邮件失败等问题请使用Chrome等现代浏览器访问！</strong>\r\n    </div>\r\n</header>\r\n<div id=\"resume\">{{.resume}}</div>\r\n<div class=\"tool-bar\" id=\"tool-bar\">\r\n    <button id=\"download\">下载到本地</button>\r\n    <button id=\"send-to-email\">发送到邮箱</button>\r\n</div>\r\n<script src=\"/assets/resume.js\"></script>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1602752662, 1602752662559110000),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1602827367, 1602827367229841400),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1602827367, 1602827367228844700),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
