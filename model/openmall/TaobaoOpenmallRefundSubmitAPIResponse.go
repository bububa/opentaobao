package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交OpenMall退款单物流 API返回值 
taobao.openmall.refund.submit

提交OpenMall退款单物流
*/
type TaobaoOpenmallRefundSubmitAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundSubmitAPIResponseModel
}

// 提交OpenMall退款单物流 成功返回结果
type TaobaoOpenmallRefundSubmitAPIResponseModel struct {
    XMLName xml.Name `xml:"openmall_refund_submit_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 提交物流单成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
}
