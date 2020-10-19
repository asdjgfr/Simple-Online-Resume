package bindata

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\r\n<html lang=\"zh-CN\">\r\n<head>\r\n    <meta charset=\"UTF-8\">\r\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n    <meta name=\"apple-mobile-web-app-capable\" content=\"yes\"/>\r\n    <link rel=\"shortcut icon\" href=\"/assets/favicon.ico\">\r\n    <title>{{.title}}</title>\r\n    <link rel=\"stylesheet\" href=\"/assets/style.css\">\r\n</head>\r\n<body>\r\n<header>\r\n    <div class=\"header\" id=\"header\">\r\n        <strong>{{.i18n.Header}}</strong>\r\n    </div>\r\n</header>\r\n<div id=\"resume\">{{.resume}}</div>\r\n<div class=\"tool-bar\" id=\"tool-bar\">\r\n    <button id=\"download\">{{.i18n.Download}}</button>\r\n    <button id=\"send-to-email\">{{.i18n.SendToEmail}}</button>\r\n</div>\r\n<script src=\"/assets/resume.js\"></script>\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1603076066, 1603076066063414900),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1603076066, 1603076066063414900),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1603076626, 1603076626328369000),
		Data:     nil,
	}}, "")
