package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse
根据安全简码查询二维码详细信息 API返回值
alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get

根据安全简码查询二维码详细信息 */
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel is 根据安全简码查询二维码详细信息 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withshort_qrcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 二维码数据
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
