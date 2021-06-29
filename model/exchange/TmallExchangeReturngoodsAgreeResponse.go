package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家确认收货 APIResponse
tmall.exchange.returngoods.agree

卖家确认收货
*/
type TmallExchangeReturngoodsAgreeAPIResponse struct {
    model.CommonResponse
    TmallExchangeReturngoodsAgreeResponse
}

type TmallExchangeReturngoodsAgreeResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_returngoods_agree_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
