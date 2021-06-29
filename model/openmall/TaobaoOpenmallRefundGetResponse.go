package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取OpenMall退款单详情 APIResponse
taobao.openmall.refund.get

获取OpenMall退款单详情
*/
type TaobaoOpenmallRefundGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundGetResponse
}

type TaobaoOpenmallRefundGetResponse struct {
    XMLName xml.Name `xml:"openmall_refund_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Refund   *TopRefundVo `json:"refund,omitempty" xml:"refund,omitempty"`

    
}
