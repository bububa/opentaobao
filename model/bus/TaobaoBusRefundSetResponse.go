package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
B2B退票申请接口 APIResponse
taobao.bus.refund.set

B2B业务支持退票
*/
type TaobaoBusRefundSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_refund_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *B2BRefundOrderRp `json:"result,omitempty" xml:"