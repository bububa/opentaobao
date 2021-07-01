package aligenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsAligenieDeviceUnbindAPIResponse
设备解绑操作接口 API返回值
alibaba.ailabs.aligenie.device.unbind

让开发者能根据设备ID进行解绑操作的接口 */
type AlibabaAilabsAligenieDeviceUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieDeviceUnbindAPIResponseModel
}

// AlibabaAilabsAligenieDeviceUnbindAPIResponseModel is 设备解绑操作接口 成功返回结果
type AlibabaAilabsAligenieDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解绑是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
