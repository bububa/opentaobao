package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝原因列表 APIResponse
taobao.refund.refusereason.get

获取商家拒绝原因列表
*/
type TaobaoRefundRefusereasonGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"refund_refusereason_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 卖家拒绝原因对象
    
    Reasons   []Reason `json:"reasons,omitempty" xml:"