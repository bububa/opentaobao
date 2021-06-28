package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝原因列表 APIResponse
taobao.refund.refusereason.get

获取商家拒绝原因列表
*/
type TaobaoRefundRefusereasonGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoRefundRefusereasonGetResponse `json:"refund_refusereason_get_response,omitempty"` 
    TaobaoRefundRefusereasonGetResponse
}

/* model for simplify = false
type TaobaoRefundRefusereasonGetResponse struct {

    // 卖家拒绝原因对象
    
    Reasons  struct {
        Reason  []Reason `json:"reason,omitempty"`
    } `json:"reasons,omitempty"`
    

    // 原因个数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty"`
    

}
*/

type TaobaoRefundRefusereasonGetResponse struct {

    // 卖家拒绝原因对象
    Reasons   []Reason `json:"reasons,omitempty"`

    // 原因个数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

}
