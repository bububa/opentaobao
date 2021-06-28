package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询活动范围 APIResponse
taobao.ump.range.get

查询某个卖家所有参加或者不参加某项活动的物品
*/
type TaobaoUmpRangeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoUmpRangeGetResponse `json:"ump_range_get_response,omitempty"` 
    TaobaoUmpRangeGetResponse
}

/* model for simplify = false
type TaobaoUmpRangeGetResponse struct {

    // 营销范围类列表！
    
    Ranges  struct {
        Range  []Range `json:"range,omitempty"`
    } `json:"ranges,omitempty"`
    

}
*/

type TaobaoUmpRangeGetResponse struct {

    // 营销范围类列表！
    Ranges   []Range `json:"ranges,omitempty"`

}
