package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
特殊部分退纠纷单查询 APIResponse
taobao.special.refund.get

获取单笔特殊部分退的纠纷单查询
*/
type TaobaoSpecialRefundGetAPIResponse struct {
    model.CommonResponse
    TaobaoSpecialRefundGetResponse
}

type TaobaoSpecialRefundGetResponse struct {
    XMLName xml.Name `xml:"special_refund_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 退款详情
    
    Refund   *Refund `json:"refund,omitempty" xml:"refund,omitempty"`

    
}
