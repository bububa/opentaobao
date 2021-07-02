package shop

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorClipbroadAPIResponse Weex页面设置或读取剪切板 API返回值
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
type AlibabaInteractSensorClipbroadAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorClipbroadAPIResponseModel
}

// AlibabaInteractSensorClipbroadAPIResponseModel is Weex页面设置或读取剪切板 成功返回结果
type AlibabaInteractSensorClipbroadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_clipbroad_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 客户端鉴权使用，实际不会发送或接收数据
	Unnamed string `json:"unnamed,omitempty" xml:"unnamed,omitempty"`
}
