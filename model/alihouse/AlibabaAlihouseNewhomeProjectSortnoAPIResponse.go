package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSortnoAPIResponse 新房排序值同步 API返回值
// alibaba.alihouse.newhome.project.sortno
//
// 新房排序值同步
type AlibabaAlihouseNewhomeProjectSortnoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectSortnoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSortnoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectSortnoAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectSortnoAPIResponseModel is 新房排序值同步 成功返回结果
type AlibabaAlihouseNewhomeProjectSortnoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_sortno_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectSortnoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSortnoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectSortnoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSortnoAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectSortnoAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectSortnoAPIResponse
func GetAlibabaAlihouseNewhomeProjectSortnoAPIResponse() *AlibabaAlihouseNewhomeProjectSortnoAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectSortnoAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectSortnoAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectSortnoAPIResponse 将 AlibabaAlihouseNewhomeProjectSortnoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectSortnoAPIResponse(v *AlibabaAlihouseNewhomeProjectSortnoAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectSortnoAPIResponse.Put(v)
}
