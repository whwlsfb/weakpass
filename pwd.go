package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var flag_username = flag.String("n", "", "请输入主要密码字段：用户名,多个字段，以空格隔开")
var flag_corp = flag.String("c", "", "请输入主要密码字段：公司名,多个字段，以空格隔开")
var flag_secondPass = flag.String("s", "", "请输入次要密码字段，自定义，多个字段，以空格隔开")
var flag_date = flag.String("d", "", "请输入时间：2019-01-01，多个字段，以空格隔开")
var flag_email = flag.String("e", "", "请输入email:xxxx@xx.com")

var sourcespass []string
var mainpass []string
var nompass []string
var bytepass = []string{"!", "@", "#", "$", "!@#"}
var count int = 0

//读文件
func readPass(filename string) string {
	passFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file error!!")
		return "nil"
	}

	return string(passFile)
}

// 写文件
func WriteToFile(filename string, content string) {
	fileObj, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		fmt.Println(content)
	}
}

func inArray(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func joinPass() {
	c := []string{"@", "#", "!", "$", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, v1 := range mainpass {
		// if len(v1) > 6 && len(v1) < 16 {
		// 	fmt.Println(v1)
		// 	count++
		// }
		if inArray(string(v1[0]), c) {
			continue
		}
		for _, v2 := range sourcespass {
			v := v1 + v2
			if v1 == v2 || len(v) > 16 {
				continue
			}
			if len(v) > 6 {
				// fmt.Println(v)
				WriteToFile("result_weak_pass.txt", v + "\n")
				count++
			}
			/*
				for _, v3 := range sourcespass {
					if v1 == v2 || v1 == v3 || v2 == v3 {
						continue
					}
					v := v1 + v2 + v3
					if len(v) > 16 {
						continue
					}
					if len(v) > 6 {
						fmt.Println(v)
						count++
					}
				}
			*/
		}
	}
}

func dealUsername(str string) {
	if strings.Contains(str, " ") {
		splitStr := strings.Split(str, " ")
		// splitStr1 := strings.Split(str, " ")
		//取每个字符串的首字母
		var shortname string
		for _, v := range splitStr {
			shortname = shortname + string(v[0])
		}

		//取姓+名字首字母
		var firstlname string
		if len(splitStr) > 2 {
			firstlname = splitStr[0] + string(splitStr[len(splitStr)-2][0]) + string(splitStr[len(splitStr)-1][0])
		}

		//名字大于三个字，取名字后两个字的拼音
		var lastname string
		if len(splitStr) > 2 {
			lastname = splitStr[len(splitStr)-2] + splitStr[len(splitStr)-1]
			splitStr = append(splitStr, lastname)
		}

		if len(splitStr) > 2 {
			splitStr = append(splitStr, shortname, strings.ToTitle(shortname), firstlname)
		} else {
			splitStr = append(splitStr, shortname, strings.ToTitle(shortname))
		}

		//拼接姓名
		allstr := strings.Replace(str, " ", "", -1)
		splitStr = append(splitStr, allstr)

		//取首字母大写
		for _, v := range splitStr {
			// fmt.Println()
			v1 := ""
			if unicode.IsLower([]rune(string(v[0]))[0]) {
				v1 = strings.ToTitle(string(v[0])) + string(v[1:])
			} else {
				v1 = strings.ToLower(string(v[0])) + string(v[1:])
			}
			splitStr = append(splitStr, v1)

		}
		fmt.Println(splitStr)
		for _, v := range splitStr {
			mainpass = append(mainpass, v)
		}
	} else {
		v1 := ""
		if unicode.IsLower([]rune(string(str))[0]) {
			v1 = strings.ToTitle(string(str[0])) + string(str[1:])
		} else {
			v1 = strings.ToLower(string(str[0])) + string(str[1:])
		}
		mainpass = append(mainpass, str, v1)
		fmt.Println(mainpass)
	}
}

func dealWithother(str string) {
	if strings.Contains(str, " ") {
		splitStr := strings.Split(str, " ")
		for _, v := range splitStr {
			sourcespass = append(sourcespass, v)
		}
	} else {
		sourcespass = append(sourcespass, str)
	}
}

func dealDate(date string) {
	if strings.Contains(date, "-") {
		datelist := strings.Split(date, "-")
		//日期
		dateday := datelist[1] + datelist[2]
		//完整日期
		alldate := strings.Replace(date, "-", "", -1)

		datelist = append(datelist, datelist[0][2:], dateday, datelist[0][2:]+dateday, alldate)

		sourcespass = append(sourcespass, datelist...)
		fmt.Println(sourcespass)
	}
}

func dealEmail(email string) {
	if strings.Contains(email, "@") {
		emailwords := strings.Split(email, "@")
		mainpass = append(mainpass, emailwords[0])
		fmt.Println(mainpass)
	}
}

//特殊字符处理
func dealByteWord(bytepass []string) {
	for _, v1 := range sourcespass {
		for _, v2 := range bytepass {
			// fmt.Println(v1 + v2)
			if v1 != "" && v2 != "" {
				sourcespass = append(sourcespass, delRN(v1)+addRN(v2), delRN(v2)+addRN(v1))
			}
		}
	}
	sourcespass = append(sourcespass, bytepass...)
}
func addRN(src string) string {
	if strings.Contains(src, "\r") {
		return src
	} else {return src + "\r"}
}
func delRN(src string) string {
	return strings.Replace(src, "\r", "", -1)
}
//处理初始密码，如123，888，520
func dealInitPass() {
	initpasslist := readPass("initpass.txt")
	passlist := strings.Split(initpasslist, "\n")
	sourcespass = append(sourcespass, passlist...)
	// fmt.Println(passlist)
}

func main() {
	// fmt.Println("hello")
	flag.Parse()
	username := string(*flag_username)
	corpname := string(*flag_corp)
	secondPass := string(*flag_secondPass)
	date := string(*flag_date)
	email := string(*flag_email)

	dealUsername(username)
	if corpname != "" {
		dealUsername(corpname)
	}

	dealDate(date)
	if email != "" {
		dealEmail(email)
	}

	dealWithother(secondPass)
	dealInitPass()
	dealByteWord(bytepass)

	fmt.Println(sourcespass)
	fmt.Println(mainpass)
	joinPass()
	weakpass := readPass("weakpass.txt")
	weakpasslist := strings.Split(weakpass, "\n")

	for _, w := range weakpasslist {
		WriteToFile("result_weak_pass.txt", w+"\n")
		count++
	}
	fmt.Printf("生成弱口令：%d个\n", count)
}
