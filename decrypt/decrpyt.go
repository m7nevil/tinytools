package decrypt

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Run() {
	now := time.Now()
	str := fmt.Sprintf("%d%d%d%d%d", now.Minute(), now.Hour(), now.Day(), int(now.Month()), now.Year())
	i, _ := strconv.Atoi(str)
	fmt.Println("您的密码将在1分钟内失效：")
	fmt.Println(strings.ToUpper(strconv.FormatInt(int64(i), 16)))
	fmt.Scanf("h")
}
