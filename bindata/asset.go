package bindata

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\n<html lang=\"zh-CN\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <meta name=\"apple-mobile-web-app-capable\" content=\"yes\"/>\n    <link rel=\"shortcut icon\" href=\"/assets/favicon.ico\">\n    <title>{{.title}}</title>\n    <link rel=\"stylesheet\" href=\"/assets/style.css\">\n    {{range $i, $v := .scripts}}\n    <script src=\"{{$v}}\" type=\"text/javascript\"></script>\n    {{end}}\n</head>\n<body>\n<header>\n    <div class=\"header\" id=\"header\">\n        <strong>{{.i18n.Header}}</strong>\n    </div>\n</header>\n<div id=\"content\">{{.content}}</div>\n<div class=\"tool-bar\" id=\"tool-bar\">\n    <button id=\"download\">{{.i18n.Download}}</button>\n    <button id=\"send-to-email\">{{.i18n.SendToEmail}}</button>\n</div>\n<script src=\"/assets/index.js\" type=\"text/javascript\"></script>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1614605058, 1614605058213344098),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1614604954, 1614604954625997978),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1614604954, 1614604954625997978),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
