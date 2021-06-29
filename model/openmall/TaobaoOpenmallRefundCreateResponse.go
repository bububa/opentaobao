package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OpenMall退款单 APIResponse
taobao.openmall.refund.create

创建OpenMall退款单
如存在未完结的退款单，则返回该退款单ID
*/
type TaobaoOpenmallRefundCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundCreateResponse
}

type TaobaoOpenmallRefundCreateResponse struct {
    XMLName xml.Name `xml:"openmall_refund_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 退款ID
    
    RefundId   int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`

    
    // 退款单状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`

    
}
