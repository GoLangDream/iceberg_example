package works

import (
	"fmt"
	"github.com/GoLangDream/iceberg/work"
)

func init() {
	work.Register("每10秒输出", "@every 10s", func() {
		fmt.Println("测试任务输出")
	})
}
