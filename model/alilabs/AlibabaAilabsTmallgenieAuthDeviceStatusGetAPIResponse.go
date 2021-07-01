package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse
设备状态查询 API返回值
alibaba.ailabs.tmallgenie.auth.device.status.get

提供给天猫精灵定制机厂商 查询设备在线状态值 */
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel is 设备状态查询 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口对象封装
	Result *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
