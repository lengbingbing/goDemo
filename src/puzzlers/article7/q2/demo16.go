package main

import (
	"github.com/tebeka/selenium"

	"github.com/tebeka/selenium/chrome"
	"log"
	"fmt"
	"time"
)


func main() {



	StartChrome()


}
func trace(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() {
		fmt.Printf("exit %s (%s)\n", msg, time.Since(start))
	}
}

// StartChrome 启动谷歌浏览器headless模式
func StartChrome() {

	defer trace("StartChrome")()
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName":                      "chrome",
	}

	// 禁止加载图片，加快渲染速度
	//imagCaps := map[string]interface{}{
	//	"profile.managed_default_content_settings.images": 2,
	//}
	//
	username := "404453996"
	password := "4qi0mx2g"

	// 代理服务器
	proxy_raw := "58.209.42.78:19586"
	proxy_str := fmt.Sprintf("http://%s:%s@%s", username, password, proxy_raw)


	log.Printf(proxy_str)

	chromeCaps := chrome.Capabilities{
		//Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式
			//"--no-sandbox",
			"--proxy="+proxy_str,
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},

	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("/usr/local/chromedriver", 9515, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	//webDriver.AddCookie(&selenium.Cookie{
	//	Name:  "defaultJumpDomain",
	//	Value: "www",
	//})
	//// 导航到目标网站
	//

	err = webDriver.Get("https://www.baidu.com")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}


	err = webDriver.Get("http://www.gd9d.com")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}


	err = webDriver.Get("https://club.m.autohome.com.cn/partner/yidian/thread/78791769")
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	webDriver.MaximizeWindow("");
	webDriver.SetImplicitWaitTimeout(5000)
	log.Println(webDriver.Title())
	log.Println(webDriver.GetCookies())
	log.Println(service)






}


