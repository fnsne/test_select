# 目的： 驗證select的先後順序
- select 有不同的channel同時可以通時，會隨機選一個case執行。
- select default的部分，只要一有任意channel case通的話，就不會去執行。
    - 所以在一些需要強制馬上關閉的情景下，可以利用這點，例如cancel的功能
```Go
package main
import (

"context"
"fmt"
"time"
)

func main(){
    // prepare data
    queue := make(chan int)
    for i := 0; i < 10; i++ {
        queue <- 1
    }
    
    ctx, cancelFunc := context.WithCancel(context.Background())
    go printQueData(ctx, queue)
    time.Sleep(10*time.Second)
    cancelFunc()
}

func printQueData(ctx context.Context, queue chan int){
    select{
    case <- ctx.Done():
        return
    default:
        // 實際要做的事，一定要放在這邊，即使要用select來接其他非context的channel，一定要分層。
        for data := range queue {
            fmt.Println(data)
        }
    }
}
```
