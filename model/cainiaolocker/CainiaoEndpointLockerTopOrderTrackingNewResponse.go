package cainiaolocker

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
事件回传接口 API返回值 
cainiao.endpoint.locker.top.order.tracking.new

用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
*/
type CainiaoEndpointLockerTopOrderTrackingNewAPIResponse struct {
    model.CommonResponse
    CainiaoEndpointLockerTopOrderTrackingNewResponse
}

// 事件回传接口 成功返回结果
type CainiaoEndpointLockerTopOrderTrackingNewResponse struct {
    XMLName xml.Name `xml:"cainiao_endpoint_locker_top_order_tracking_new_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
