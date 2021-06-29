package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIResponse
taobao.openmall.trade.create

创建Openmall订单
*/
type TaobaoOpenmallTradeCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeCreateResponse
}

type TaobaoOpenmallTradeCreateResponse struct {
    XMLName xml.Name `xml:"openmall_trade_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
