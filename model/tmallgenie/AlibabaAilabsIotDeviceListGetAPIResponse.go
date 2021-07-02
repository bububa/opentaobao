package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListGetAPIResponse 获取iot设备列表 API返回值
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
type AlibabaAilabsIotDeviceListGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsIotDeviceListGetAPIResponseModel
}

// AlibabaAilabsIotDeviceListGetAPIResponseModel is 获取iot设备列表 成功返回结果
type AlibabaAilabsIotDeviceListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_iot_device_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsIotDeviceListGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
