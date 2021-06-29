package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意换货申请 APIResponse
tmall.exchange.agree

卖家同意换货申请
*/
type TmallExchangeAgreeAPIResponse struct {
    model.CommonResponse
    TmallExchangeAgreeResponse
}

type TmallExchangeAgreeResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_agree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
