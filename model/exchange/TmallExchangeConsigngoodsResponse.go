package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家发货 APIResponse
tmall.exchange.consigngoods

卖家发货
*/
type TmallExchangeConsigngoodsAPIResponse struct {
    model.CommonResponse
    TmallExchangeConsigngoodsResponse
}

type TmallExchangeConsigngoodsResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_consigngoods_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *RefundBaseResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
