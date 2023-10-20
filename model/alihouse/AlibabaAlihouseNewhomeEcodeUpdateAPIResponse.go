package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeEcodeUpdateAPIResponse 新房货变更E码 API返回值
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
type AlibabaAlihouseNewhomeEcodeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeEcodeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel is 新房货变更E码 成功返回结果
type AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_ecode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeEcodeUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeEcodeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeEcodeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeEcodeUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeEcodeUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeEcodeUpdateAPIResponse
func GetAlibabaAlihouseNewhomeEcodeUpdateAPIResponse() *AlibabaAlihouseNewhomeEcodeUpdateAPIResponse {
	return poolAlibabaAlihouseNewhomeEcodeUpdateAPIResponse.Get().(*AlibabaAlihouseNewhomeEcodeUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeEcodeUpdateAPIResponse 将 AlibabaAlihouseNewhomeEcodeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeEcodeUpdateAPIResponse(v *AlibabaAlihouseNewhomeEcodeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeEcodeUpdateAPIResponse.Put(v)
}
