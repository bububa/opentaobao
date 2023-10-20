package damaiticklet

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMxOpengatewayScriptAPIResponse 第三方剧本数据推送 API返回值
// alibaba.damai.mx.opengateway.script
//
// 第三方剧本数据推送
type AlibabaDamaiMxOpengatewayScriptAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMxOpengatewayScriptAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMxOpengatewayScriptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMxOpengatewayScriptAPIResponseModel).Reset()
}

// AlibabaDamaiMxOpengatewayScriptAPIResponseModel is 第三方剧本数据推送 成功返回结果
type AlibabaDamaiMxOpengatewayScriptAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mx_opengateway_script_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值模型
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMxOpengatewayScriptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMxOpengatewayScriptAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMxOpengatewayScriptAPIResponse)
	},
}

// GetAlibabaDamaiMxOpengatewayScriptAPIResponse 从 sync.Pool 获取 AlibabaDamaiMxOpengatewayScriptAPIResponse
func GetAlibabaDamaiMxOpengatewayScriptAPIResponse() *AlibabaDamaiMxOpengatewayScriptAPIResponse {
	return poolAlibabaDamaiMxOpengatewayScriptAPIResponse.Get().(*AlibabaDamaiMxOpengatewayScriptAPIResponse)
}

// ReleaseAlibabaDamaiMxOpengatewayScriptAPIResponse 将 AlibabaDamaiMxOpengatewayScriptAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMxOpengatewayScriptAPIResponse(v *AlibabaDamaiMxOpengatewayScriptAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMxOpengatewayScriptAPIResponse.Put(v)
}
