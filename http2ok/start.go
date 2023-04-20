package http2ok

import (
	"fmt"
	"http2ok/common"
	"sync"
)

var successlist []common.SuccessTarget

func Start() {
	if common.URL == "" && common.ICP == "" && common.ICPFile == "" && common.TargetFile == "" {
		fmt.Println("Usage: http2ok.exe [-u|-icp|-tf|-if|-h] [values].")
		return
	}

	if common.Checkproxy {
		checkProxy(common.Proxy)
	}

	if common.ICP != "" {
		//域名备案查询
		keyword, flag := CheckDomain(common.ICP)
		if flag {
			_, _ = ICPSearch(keyword)
		}
		return
	} else if common.ICPFile != "" {
		domains, err := common.Readfile(common.ICPFile)
		if err != nil {
			common.LogError(err)
		}
		for _, domain := range domains {
			//域名备案查询
			keyword, flag := CheckDomain(domain)
			if flag {
				_, _ = ICPSearch(keyword)
			}
		}
		return
	}

	//获取urlList
	urlList := []string{}
	if common.URL != "" {
		urlList = append(urlList, common.URL)
	}

	if common.TargetFile != "" {
		hosts, err := common.Readfile(common.TargetFile)
		if err != nil {
			common.LogError(err)
		}
		for _, host := range hosts {
			urlList = append(urlList, host)
		}
	}

	common.Num = len(urlList)

	workers := common.Threads
	urls := make(chan string, common.Num)
	results := make(chan common.SuccessTarget, common.Num)
	var wg sync.WaitGroup

	//接收结果
	go func() {
		for found := range results {
			successlist = append(successlist, found)
			common.SuccessNum += 1
			wg.Done()
		}
	}()

	//多线程扫描
	for i := 0; i < workers; i++ {
		go func() {
			for u := range urls {
				GetWeb(u, results, &wg)
				wg.Done()
			}
		}()
	}

	//添加扫描任务
	for _, url := range urlList {
		wg.Add(1)
		urls <- url
	}
	wg.Wait()
	close(urls)
	close(results)

	//导出成功的url
	if len(successlist) != 0 {
		fmt.Println("成功结果如下：")
		for _, t := range successlist {
			result := fmt.Sprintf("%s", t.Url)
			common.WriteFile(result, "okurl.txt")
			fmt.Printf("%s\t%s\t%d\t%s\t%s\t%s\n", "[+] ", t.Url, t.Code, t.Title, t.ICPServiceLicence, t.ICPUnitName)
		}
	}
}
