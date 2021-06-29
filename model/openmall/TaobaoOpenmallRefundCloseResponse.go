package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭OpenMall退款单 APIResponse
taobao.openmall.refund.close

关闭OpenMall退款单
*/
type TaobaoOpenmallRefundCloseAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundCloseResponse
}

type TaobaoOpenmallRefundCloseResponse struct {
    XMLName xml.Name `xml:"openmall_refund_close_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否关闭成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
