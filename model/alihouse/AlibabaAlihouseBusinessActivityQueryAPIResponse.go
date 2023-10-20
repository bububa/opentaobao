package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseBusinessActivityQueryAPIResponse 电商券活动公司数据查询 API返回值
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
type AlibabaAlihouseBusinessActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseBusinessActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseBusinessActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseBusinessActivityQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseBusinessActivityQueryAPIResponseModel is 电商券活动公司数据查询 成功返回结果
type AlibabaAlihouseBusinessActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_business_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseBusinessActivityQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseBusinessActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseBusinessActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseBusinessActivityQueryAPIResponse)
	},
}

// GetAlibabaAlihouseBusinessActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseBusinessActivityQueryAPIResponse
func GetAlibabaAlihouseBusinessActivityQueryAPIResponse() *AlibabaAlihouseBusinessActivityQueryAPIResponse {
	return poolAlibabaAlihouseBusinessActivityQueryAPIResponse.Get().(*AlibabaAlihouseBusinessActivityQueryAPIResponse)
}

// ReleaseAlibabaAlihouseBusinessActivityQueryAPIResponse 将 AlibabaAlihouseBusinessActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseBusinessActivityQueryAPIResponse(v *AlibabaAlihouseBusinessActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseBusinessActivityQueryAPIResponse.Put(v)
}
