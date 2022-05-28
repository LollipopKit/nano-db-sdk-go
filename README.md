## Nano DB SDK for Go

### 使用方法
```go
import (
    ndb "git.lolli.tech/lollipopkit/nano-db-sdk-go"
)
var (
    // 创建一个数据库对象
    db = ndb.NewDB("DB Url", "DB Cookie")
)

type Chapter struct {
    Id int64
    Cid int64
    Content string
    Title string
}

func main() {
    // 写入数据
    var chapter1 Chapter
    chapter1.Title = "第一章"
    err = db.Write("novel/bookname/1.json", chapter1)
    if err != nil {
        panic(err)
    }

    // 读取数据
    err := db.Read("novel/bookname/1.json", &chapter1)
    if err != nil {
        panic(err)
    }
    println(chapter1.Title)

    // 删除数据
    err = db.Delete("novel/bookname/1.json")
    if err != nil {
        panic(err)
    }
}
```

**剩余方法**请查看源码