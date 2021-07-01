package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotDeviceIsexistAPIResponse
判断设备是否存在 API返回值
alibaba.aliqin.fc.iot.device.isexist

判断设备是否存在 */
type AlibabaAliqinFcIotDeviceIsexistAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotDeviceIsexistAPIResponseModel
}

// AlibabaAliqinFcIotDeviceIsexistAPIResponseModel is 判断设备是否存在 成功返回结果
type AlibabaAliqinFcIotDeviceIsexistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_isexist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotDeviceIsexistResult `json:"result,omitempty" xml:"result,omitempty"`
}
