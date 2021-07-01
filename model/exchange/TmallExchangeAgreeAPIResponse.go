package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意换货申请 API返回值 
tmall.exchange.agree

卖家同意换货申请
*/
type TmallExchangeAgreeAPIResponse struct {
    model.CommonResponse
    TmallExchangeAgreeAPIResponseModel
}

// 卖家同意换货申请 成功返回结果
type TmallExchangeAgreeAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_exchange_agree_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
