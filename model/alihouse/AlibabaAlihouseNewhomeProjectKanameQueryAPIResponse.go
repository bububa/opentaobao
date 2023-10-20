package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse 查询KA楼盘名称 API返回值
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
type AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectKanameQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectKanameQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectKanameQueryAPIResponseModel is 查询KA楼盘名称 成功返回结果
type AlibabaAlihouseNewhomeProjectKanameQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_kaname_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectKanameQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectKanameQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse
func GetAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse() *AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse 将 AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse(v *AlibabaAlihouseNewhomeProjectKanameQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectKanameQueryAPIResponse.Put(v)
}
