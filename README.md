# pinduoduo-sdk
拼多多的golang版本的sdk

```
    package main
    
    import (
    	"fmt"
    
    	"github.com/xiangwork/pinduoduo-sdk/common"
    	"github.com/xiangwork/pinduoduo-sdk/ddk"
    )
    
    func main() {
    	api := ddk.NewDuoduoKe(common.NewContext("client_id", "client_secret"))
    
    	result, err := api.GoodsSearch(nil)
    	if err != nil {
    		fmt.Println(err)
    		return
    	}
    	fmt.Println(result)
    }
 
```

###

觉得对你有帮助,给个start. 也欢迎大家一起完善这个sdk