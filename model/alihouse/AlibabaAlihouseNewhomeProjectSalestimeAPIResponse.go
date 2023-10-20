package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectSalestimeAPIResponse 楼盘销售时刻同步 API返回值
// alibaba.alihouse.newhome.project.salestime
//
// 楼盘销售时刻同步
type AlibabaAlihouseNewhomeProjectSalestimeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectSalestimeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSalestimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectSalestimeAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectSalestimeAPIResponseModel is 楼盘销售时刻同步 成功返回结果
type AlibabaAlihouseNewhomeProjectSalestimeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_salestime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectSalestimeResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectSalestimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectSalestimeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectSalestimeAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectSalestimeAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectSalestimeAPIResponse
func GetAlibabaAlihouseNewhomeProjectSalestimeAPIResponse() *AlibabaAlihouseNewhomeProjectSalestimeAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectSalestimeAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectSalestimeAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectSalestimeAPIResponse 将 AlibabaAlihouseNewhomeProjectSalestimeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectSalestimeAPIResponse(v *AlibabaAlihouseNewhomeProjectSalestimeAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectSalestimeAPIResponse.Put(v)
}
