package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改OpenMall退款申请 APIResponse
taobao.openmall.refund.modify

修改OpenMall退款申请
*/
type TaobaoOpenmallRefundModifyAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundModifyResponse
}

type TaobaoOpenmallRefundModifyResponse struct {
    XMLName xml.Name `xml:"openmall_refund_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
