package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际订单清关信息 APIResponse
tmall.hk.order.clearance.query

天猫国际订单清关信息查询
*/
type TmallHkOrderClearanceQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_hk_order_clearance_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果码
    
    ResponseCode   string `json:"response_code,omitempty" xml:"