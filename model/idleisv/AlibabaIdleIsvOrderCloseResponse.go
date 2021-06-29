package idleisv

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商闲鱼卖家主动关闭订单 API返回值 
alibaba.idle.isv.order.close

供外部服务商 isv 提供卖家主动关闭交易订单的功能
*/
type AlibabaIdleIsvOrderCloseAPIResponse struct {
    model.CommonResponse
    AlibabaIdleIsvOrderCloseResponse
}

// 服务商闲鱼卖家主动关闭订单 成功返回结果
type AlibabaIdleIsvOrderCloseResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_isv_order_close_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回结果
    Result   *AlibabaIdleIsvOrderCloseResult `json:"result,omitempty" xml:"result,omitempty"`
}
