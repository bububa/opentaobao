package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际订单清关信息 APIResponse
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
type TmallHkOrderClearanceQueryAPIResponse struct {
    model.CommonResponse
    // Response *TmallHkOrderClearanceQueryResponse `json:"tmall_hk_order_clearance_query_response,omitempty"` 
    TmallHkOrderClearanceQueryResponse
}

/* model for simplify = false
type TmallHkOrderClearanceQueryResponse struct {

    // 结果码
    
    ResponseCode   string `json:"response_code,omitempty"`
    

    // 结果描述
    
    ResponseMessage   string `json:"response_message,omitempty"`
    

    // 清关数据封装
    
    Obj  *struct {
        ClearanceDataDo  *ClearanceDataDo `json:"clearance_data_do,omitempty"`
    } `json:"obj,omitempty"`
    

    // 是否正常
    
    Succeeded   bool `json:"succeeded,omitempty"`
    

}
*/

type TmallHkOrderClearanceQueryResponse struct {

    // 结果码
    ResponseCode   string `json:"response_code,omitempty"`

    // 结果描述
    ResponseMessage   string `json:"response_message,omitempty"`

    // 清关数据封装
    Obj   *ClearanceDataDo `json:"obj,omitempty"`

    // 是否正常
    Succeeded   bool `json:"succeeded,omitempty"`

}
