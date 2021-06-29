package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API返回值 
taobao.openmall.trade.create

创建Openmall订单
*/
type TaobaoOpenmallTradeCreateAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeCreateResponse
}

// 创建订单 成功返回结果
type TaobaoOpenmallTradeCreateResponse struct {
    XMLName xml.Name `xml:"openmall_trade_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
