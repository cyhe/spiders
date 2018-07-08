package regex

import (
	"regexp"
	"fmt"
)

const text = "My email is sunrisesoar@gmail.com.cn.org email1 is aaa@bbb.com email2 is qqq@qqq.com.cn"

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
