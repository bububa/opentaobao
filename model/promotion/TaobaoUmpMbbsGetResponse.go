package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取营销积木块列表 APIResponse
taobao.ump.mbbs.get

获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的
*/
type TaobaoUmpMbbsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpMbbsGetResponse `json:"ump_mbbs_get_response,omitempty"` 
    TaobaoUmpMbbsGetResponse
}

/* model for simplify = false
type TaobaoUmpMbbsGetResponse struct {

    // 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    
    Mbbs  struct {
        String  []string `json:"string,omitempty"`
    } `json:"mbbs,omitempty"`
    

}
*/

type TaobaoUmpMbbsGetResponse struct {

    // 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    Mbbs   []string `json:"mbbs,omitempty"`

}
