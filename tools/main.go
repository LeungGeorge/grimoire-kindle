package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// README TODO
var README = `
# 魔法书之电子书篇


- [亚马逊官网电子书](https://www.amazon.cn/Kindle%E7%94%B5%E5%AD%90%E4%B9%A6/b?ie=UTF8&node=116169071&ref_=nav_topnav_giftcert)    
- [Kindle电子书杂志分享的博客](http://blog.sina.com.cn/u/6559127255)     
- [读书达人](http://www.dushudaren.com/)


电子书
https://b-ok.cc/
好的地方在于可以找paper
www.jiumodiary.com

部分书可以从这上面找，中英文都有     

http://gen.lib.rus.ec/ 

https://manybooks.net/  

https://www.librarything.com/ 

以上三个网站可找英文书

kindle免费电子书资源
http://www.58wzb.com/page/2/?s=%E6%96%87%E5%8C%96


一键发送到Kindle - 选项
chrome-extension://ipkfnchcgalnafehpglfbommidgmalan/options.html

知乎日报-rss
http://news-at.zhihu.com/api/4/stories/latest?client=0


kindle订阅杂志
rss精选

[Kindle 相关工具](https://bookfere.com/tools#ClippingsFere)

## 目录

`

func main() {
	baseDir, _ := os.Getwd()
	arrDir := strings.Split(baseDir, "/")
	newBaseDir := strings.Join(arrDir[0:len(arrDir)-1], "/")
	filepath.Walk(newBaseDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		note := strings.TrimPrefix(path, baseDir)
		if note == "" {
			return nil
		}
		notePath := strings.TrimPrefix(path, newBaseDir)
		if strings.HasPrefix(notePath, "/.git/") ||
			strings.HasPrefix(notePath, "/tools/") ||
			strings.HasPrefix(notePath, "/.gitignore") ||
			strings.HasPrefix(notePath, "/LICENSE") ||
			strings.HasSuffix(notePath, "README.md") ||
			strings.HasSuffix(notePath, ".DS_Store") {
			return nil
		}

		markdownList := fmt.Sprintf("- [%v](%v)\n", notePath, url.QueryEscape(notePath))
		README += markdownList
		return nil
	})

	fmt.Println(README)
	os.WriteFile("../README.md", []byte(README), os.ModePerm)
}
