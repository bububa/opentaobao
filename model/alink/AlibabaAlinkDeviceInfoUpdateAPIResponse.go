package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceInfoUpdateAPIResponse 更新设备昵称等信息 API返回值
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
type AlibabaAlinkDeviceInfoUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceInfoUpdateAPIResponseModel
}

// AlibabaAlinkDeviceInfoUpdateAPIResponseModel is 更新设备昵称等信息 成功返回结果
type AlibabaAlinkDeviceInfoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_info_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
