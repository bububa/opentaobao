package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购获取订单列表 API返回值 
taobao.opentrade.group.order

组团购场景下，获取开团的订单列表
*/
type TaobaoOpentradeGroupOrderAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupOrderAPIResponseModel
}

// 组团购获取订单列表 成功返回结果
type TaobaoOpentradeGroupOrderAPIResponseModel struct {
    XMLName xml.Name `xml:"opentrade_group_order_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 订单id列表
    Result   []string `json:"result,omitempty" xml:"result>string,omitempty"`
}
