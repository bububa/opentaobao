package larkiot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道卖品落单 API返回值 
taobao.lark.iot.order.confirmorder

云智对接无人超市，接收无人超市订单信息
*/
type TaobaoLarkIotOrderConfirmorderAPIResponse struct {
    model.CommonResponse
    TaobaoLarkIotOrderConfirmorderAPIResponseModel
}

// iot渠道卖品落单 成功返回结果
type TaobaoLarkIotOrderConfirmorderAPIResponseModel struct {
    XMLName xml.Name `xml:"lark_iot_order_confirmorder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回下单结果
    Data   *BizSingleResult `json:"data,omitempty" xml:"data,omitempty"`
}
