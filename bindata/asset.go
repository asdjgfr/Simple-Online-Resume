package bindata

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!DOCTYPE html>\n<html lang=\"zh-CN\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <meta name=\"apple-mobile-web-app-capable\" content=\"yes\"/>\n    <link rel=\"shortcut icon\" href=\"/assets/favicon.ico\">\n    <title>{{.title}}</title>\n    <link rel=\"stylesheet\" href=\"/assets/style.css\">\n</head>\n<body>\n<header>\n    <div class=\"header\" id=\"header\">\n        <strong>{{.i18n.Header}}</strong>\n    </div>\n</header>\n<div id=\"resume\">{{.resume}}</div>\n<div class=\"tool-bar\" id=\"tool-bar\">\n    <button id=\"download\">{{.i18n.Download}}</button>\n    <button id=\"send-to-email\">{{.i18n.SendToEmail}}</button>\n</div>\n<script src=\"/assets/resume.js\"></script>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1614048724, 1614048724611105700),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1603077881, 1603077881000000000),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1ff,
		Mtime:    time.Unix(1603077881, 1603077881000000000),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
