package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过ids列表获取营销积木块列表 APIResponse
taobao.ump.mbbs.list.get

通过营销积木id列表来获取营销积木块列表。
*/
type TaobaoUmpMbbsListGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpMbbsListGetResponse `json:"ump_mbbs_list_get_response,omitempty"` 
    TaobaoUmpMbbsListGetResponse
}

/* model for simplify = false
type TaobaoUmpMbbsListGetResponse struct {

    // 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    
    Mbbs  struct {
        String  []string `json:"string,omitempty"`
    } `json:"mbbs,omitempty"`
    

}
*/

type TaobaoUmpMbbsListGetResponse struct {

    // 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
    Mbbs   []string `json:"mbbs,omitempty"`

}
