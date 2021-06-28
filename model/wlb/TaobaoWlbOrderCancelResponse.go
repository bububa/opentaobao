package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消物流宝订单 APIResponse
taobao.wlb.order.cancel

取消物流宝订单
*/
type TaobaoWlbOrderCancelAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbOrderCancelResponse `json:"wlb_order_cancel_response,omitempty"` 
    TaobaoWlbOrderCancelResponse
}

/* model for simplify = false
type TaobaoWlbOrderCancelResponse struct {

    // 修改时间，只有在取消成功的情况下，才可以做
    
    ModifyTime   string `json:"modify_time,omitempty"`
    

    // 错误编码列表
    
    ErrorCodeList   string `json:"error_code_list,omitempty"`
    

}
*/

type TaobaoWlbOrderCancelResponse struct {

    // 修改时间，只有在取消成功的情况下，才可以做
    ModifyTime   string `json:"modify_time,omitempty"`

    // 错误编码列表
    ErrorCodeList   string `json:"error_code_list,omitempty"`

}
