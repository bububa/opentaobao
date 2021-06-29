package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭订单 APIResponse
taobao.openmall.trade.close

关闭订单
*/
type TaobaoOpenmallTradeCloseAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeCloseResponse
}

type TaobaoOpenmallTradeCloseResponse struct {
    XMLName xml.Name `xml:"openmall_trade_close_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 出参
    
    Result   *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
