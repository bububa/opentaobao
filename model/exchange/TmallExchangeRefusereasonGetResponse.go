package exchange

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取拒绝换货原因列表 APIResponse
tmall.exchange.refusereason.get

获取拒绝换货原因列表
*/
type TmallExchangeRefusereasonGetAPIResponse struct {
    model.CommonResponse
    TmallExchangeRefusereasonGetResponse
}

type TmallExchangeRefusereasonGetResponse struct {
    XMLName xml.Name `xml:"tmall_exchange_refusereason_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TmallExchangeRefusereasonGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
