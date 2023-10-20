package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse 置业顾问批量绑定楼盘 API返回值
// alibaba.alihouse.newhome.project.adviser.bind
//
// 置业顾问批量绑定楼盘
type AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel is 置业顾问批量绑定楼盘 成功返回结果
type AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_adviser_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectAdviserBindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectAdviserBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse
func GetAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse() *AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse 将 AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse(v *AlibabaAlihouseNewhomeProjectAdviserBindAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectAdviserBindAPIResponse.Put(v)
}
