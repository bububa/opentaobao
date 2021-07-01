package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsTmallgenieAuthGettokenAPIResponse
设备授权 API返回值
alibaba.ailabs.tmallgenie.auth.gettoken

获取设备授权码 */
type AlibabaAilabsTmallgenieAuthGettokenAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthGettokenAPIResponseModel
}

// AlibabaAilabsTmallgenieAuthGettokenAPIResponseModel is 设备授权 成功返回结果
type AlibabaAilabsTmallgenieAuthGettokenAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_gettoken_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备注册结果
	RegisterResult *RegisterInfoVo `json:"register_result,omitempty" xml:"register_result,omitempty"`
	// 授权结果
	AuthResult *DeviceTokenVo `json:"auth_result,omitempty" xml:"auth_result,omitempty"`
}
