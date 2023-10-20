package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse 提交预售证 API返回值
// alibaba.alihouse.newhome.project.presalepermit.submit
//
// 提交楼盘预售证信息
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel is 提交预售证 成功返回结果
type AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaAlihouseNewhomeProjectPresalepermitSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse
func GetAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse() *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse 将 AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse(v *AlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectPresalepermitSubmitAPIResponse.Put(v)
}
