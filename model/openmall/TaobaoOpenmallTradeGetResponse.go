package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询订单详情 APIResponse
taobao.openmall.trade.get

查询订单详情
*/
type TaobaoOpenmallTradeGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeGetResponse
}

type TaobaoOpenmallTradeGetResponse struct {
    XMLName xml.Name `xml:"openmall_trade_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TopTradeDetailVo `json:"result,omitempty" xml:"result,omitempty"`

    
}
