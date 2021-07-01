package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotDevicePostAPIResponse
商家提交设备信息 API返回值
alibaba.aliqin.fc.iot.device.post

物联网商家设备信息录入 */
type AlibabaAliqinFcIotDevicePostAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotDevicePostAPIResponseModel
}

// AlibabaAliqinFcIotDevicePostAPIResponseModel is 商家提交设备信息 成功返回结果
type AlibabaAliqinFcIotDevicePostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotDevicePostResult `json:"result,omitempty" xml:"result,omitempty"`
}
