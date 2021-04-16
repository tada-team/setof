# Sets

```go
import "setof"

func main()  {
    s := setof.Strings("aaa", "bbb", "aaa")
    print(len(s)) // 2
    s.Update("ccc", "444")
    print(s.Contains("123")) // false
    print(s.Contains("aaa")) // true
}
```
