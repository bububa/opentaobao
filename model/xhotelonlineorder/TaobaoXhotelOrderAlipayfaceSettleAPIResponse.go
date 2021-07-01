package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信用住订单结账接口 API返回值 
taobao.xhotel.order.alipayface.settle

用于离店付订单在客人离店后，发起结账以及扣款等后续动作
*/
type TaobaoXhotelOrderAlipayfaceSettleAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel
}

// 信用住订单结账接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceSettleAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_order_alipayface_settle_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
