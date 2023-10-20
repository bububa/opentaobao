package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse 新房户型变更E码 API返回值
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
type AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponseModel is 新房户型变更E码 成功返回结果
type AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_layout_ecode_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeLayoutEcodeUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse
func GetAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse() *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse {
	return poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse.Get().(*AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse 将 AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse(v *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIResponse.Put(v)
}
