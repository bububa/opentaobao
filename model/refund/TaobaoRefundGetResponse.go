package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔退款详情 APIResponse
taobao.refund.get

获取单笔退款详情
*/
type TaobaoRefundGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"refund_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款详情
    
    Refund   *Refund `json:"refund,omitempty" xml:"