package main

import (
    "fmt"
    "strconv"
    "strings"
)

func subdomainVisits(cpdomains []string) []string {
    //定义一个map存放域名对应的访问次数
    mapData := make(map[string]int64)
    for _, v := range cpdomains {
        //对数字进行切片
        count := strings.Index(v, " ")
        num, _ := strconv.ParseInt(v[:count], 10, 64)

        //对后面的字符串进行切片
		domainName := v[count+1:]
         mapData[domainName] += num
		for {
			pos := strings.Index(domainName, ".")
			if pos != -1 {
				domainName = domainName[pos+1:]
				mapData[domainName] += num
			}else{
				break
			}
		}
    }

    var res []string
    for key, value := range mapData {
		res = append(res, strconv.FormatInt(value, 10) + " " + key)
    }
    return res
}


func main() {

    arrString1 := make([]string, 1)
    arrString1[0] = "9001 discuss.leetcode.com"

    arrString2 := make([]string, 4)
    arrString2[0] = "900 google.mail.com"
    arrString2[1] = "50 yahoo.com"
    arrString2[2] = "1 intel.mail.com"
    arrString2[3] = "5 wiki.org"

    arrString3 := make([]string, 10)
    arrString3 = subdomainVisits(arrString1)

    arrString4 := make([]string, 10)
    arrString4 = subdomainVisits(arrString2)

    fmt.Println(arrString3)

    fmt.Println(arrString4)

}
