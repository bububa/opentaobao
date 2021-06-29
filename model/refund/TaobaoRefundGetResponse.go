package refund

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔退款详情 API返回值 
taobao.refund.get

获取单笔退款详情
*/
type TaobaoRefundGetAPIResponse struct {
    model.CommonResponse
    TaobaoRefundGetResponse
}

// 获取单笔退款详情 成功返回结果
type TaobaoRefundGetResponse struct {
    XMLName xml.Name `xml:"refund_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退款详情
    Refund   *Refund `json:"refund,omitempty" xml:"refund,omitempty"`
}
