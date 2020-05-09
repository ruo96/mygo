package function


/**
  1. 关键字 defer 用于注册延迟调用。
  2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
  3. 多个defer语句，按先进后出的方式执行。
  4. defer语句中的变量，在defer声明时就决定了。

defer的用处
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放

defer 是先进后出
这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。
 */

import "fmt"

func maindef1() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}
