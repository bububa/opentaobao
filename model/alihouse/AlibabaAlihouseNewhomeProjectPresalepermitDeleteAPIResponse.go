package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse 删除楼盘预售证 API返回值
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel is 删除楼盘预售证 成功返回结果
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_presalepermit_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectPresalepermitDeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse
func GetAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse 将 AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse(v *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIResponse.Put(v)
}
