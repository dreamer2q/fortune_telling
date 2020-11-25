# fortune telling

使用Go爬取了 [关帝灵签](https://www.buyiju.com/guandi/), 封装成接口, 方便调用。

> 关帝俗称关公，是世人崇拜的英雄，为保护人民的神祗，也是做生意之人必须信奉的武财神，
> 保证商品童叟无欺，跟土地公一样保佑生意兴隆、招财进宝，香火兴盛、历久不衰。
> 关帝灵签是一个很古老的占卜项目，卜易居算命网·关公灵签的签词特选取了现存通行于闽南及台湾的泉州通淮关岳庙诗签版本，
> 系清·光绪间泉州通淮关岳庙印制的秘本《关帝灵签》。

## 求签

> 求签，指迷信的人在神佛面前抽签来占卜吉凶。
> 是中国的民间习俗，是占卜的其中一种形式。现今的道观、寺庙和民间的庙宇，大多摆上签筒供人抽取签条问卜。

## Get Started

 - 引入包
 
```go
import telling "github.com/dreamer2q/fortune_telling"
```

 - example
 
```go
package main

import (
    "fmt"
    telling "github.com/dreamer2q/fortune_telling"
)

func main(){
    randKey := "you-sign-key"
    if telling.HasAsked(randKey) {
        fmt.Printf("%s has asked, come back later",randKey)
    }else{
        tell,_ := telling.Ask(randKey)
        fmt.Printf("%v",tell)
    }
}
```

 - 重置签到数据 
 
```go
telling.Reset()
```

## License

[mit](LICENSE)