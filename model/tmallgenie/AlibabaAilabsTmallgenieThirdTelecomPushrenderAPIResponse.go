package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse 电信-推送拉起设备应用 API返回值
// alibaba.ailabs.tmallgenie.third.telecom.pushrender
//
// 电信-推送拉起设备应用
type AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel is 电信-推送拉起设备应用 成功返回结果
type AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_third_telecom_pushrender_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *AlibabaAilabsTmallgenieThirdTelecomPushrenderResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse
func GetAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse() *AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse {
	return poolAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse.Get().(*AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse 将 AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse(v *AlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieThirdTelecomPushrenderAPIResponse.Put(v)
}
