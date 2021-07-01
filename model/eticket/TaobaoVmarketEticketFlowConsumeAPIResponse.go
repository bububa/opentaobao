package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无交易类凭证核销 API返回值 
taobao.vmarket.eticket.flow.consume

无交易类凭证核销
*/
type TaobaoVmarketEticketFlowConsumeAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketFlowConsumeAPIResponseModel
}

// 无交易类凭证核销 成功返回结果
type TaobaoVmarketEticketFlowConsumeAPIResponseModel struct {
    XMLName xml.Name `xml:"vmarket_eticket_flow_consume_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 执行成功
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
    // 错误提示信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
