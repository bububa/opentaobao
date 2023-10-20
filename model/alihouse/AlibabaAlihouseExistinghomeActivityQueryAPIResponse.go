package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeActivityQueryAPIResponse 五福活动经纪人获取 API返回值
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
type AlibabaAlihouseExistinghomeActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel is 五福活动经纪人获取 成功返回结果
type AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihouseExistinghomeActivityQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseExistinghomeActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeActivityQueryAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeActivityQueryAPIResponse
func GetAlibabaAlihouseExistinghomeActivityQueryAPIResponse() *AlibabaAlihouseExistinghomeActivityQueryAPIResponse {
	return poolAlibabaAlihouseExistinghomeActivityQueryAPIResponse.Get().(*AlibabaAlihouseExistinghomeActivityQueryAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeActivityQueryAPIResponse 将 AlibabaAlihouseExistinghomeActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeActivityQueryAPIResponse(v *AlibabaAlihouseExistinghomeActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeActivityQueryAPIResponse.Put(v)
}
