package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家拒绝换货申请 APIResponse
tmall.exchange.refuse

卖家拒绝换货申请
*/
type TmallExchangeRefuseAPIResponse struct {
    model.CommonResponse
    TmallExchangeRefuseResponse
}

type TmallExchangeRefuseResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_refuse_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
