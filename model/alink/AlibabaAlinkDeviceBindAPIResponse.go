package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceBindAPIResponse 绑定设备 API返回值
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
type AlibabaAlinkDeviceBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceBindAPIResponseModel
}

// AlibabaAlinkDeviceBindAPIResponseModel is 绑定设备 成功返回结果
type AlibabaAlinkDeviceBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
