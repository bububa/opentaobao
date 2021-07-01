package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceListAPIResponse
获取用户设备列表 API返回值
alibaba.ailabs.tmallgenie.auth.device.list

通过此接口获取用户绑定的设备信息列表 */
type AlibabaAilabsTmallgenieAuthDeviceListAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceListAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceListAPIResponseModel is 获取用户设备列表 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备列表
	Devices []AlibabaAilabsTmallgenieAuthDeviceListResult `json:"devices,omitempty" xml:"devices>alibaba_ailabs_tmallgenie_auth_device_list_result,omitempty"`
}
