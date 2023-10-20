package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthUicUserinfoHealthidGetAPIResponse 获取健康id API返回值
// alibaba.alihealth.uic.userinfo.healthid.get
//
// 根据支付宝用户ID获取用户健康ID
type AlibabaAlihealthUicUserinfoHealthidGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthUicUserinfoHealthidGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthUicUserinfoHealthidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthUicUserinfoHealthidGetAPIResponseModel).Reset()
}

// AlibabaAlihealthUicUserinfoHealthidGetAPIResponseModel is 获取健康id 成功返回结果
type AlibabaAlihealthUicUserinfoHealthidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_uic_userinfo_healthid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthUicUserinfoHealthidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthUicUserinfoHealthidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthUicUserinfoHealthidGetAPIResponse)
	},
}

// GetAlibabaAlihealthUicUserinfoHealthidGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthUicUserinfoHealthidGetAPIResponse
func GetAlibabaAlihealthUicUserinfoHealthidGetAPIResponse() *AlibabaAlihealthUicUserinfoHealthidGetAPIResponse {
	return poolAlibabaAlihealthUicUserinfoHealthidGetAPIResponse.Get().(*AlibabaAlihealthUicUserinfoHealthidGetAPIResponse)
}

// ReleaseAlibabaAlihealthUicUserinfoHealthidGetAPIResponse 将 AlibabaAlihealthUicUserinfoHealthidGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthUicUserinfoHealthidGetAPIResponse(v *AlibabaAlihealthUicUserinfoHealthidGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthUicUserinfoHealthidGetAPIResponse.Put(v)
}
