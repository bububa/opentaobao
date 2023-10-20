package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvUserQueryAPIResponse 服务商ISV闲鱼用户信息查询 API返回值
// alibaba.idle.isv.user.query
//
// 服务商ISV闲鱼用户信息查询
type AlibabaIdleIsvUserQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvUserQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvUserQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvUserQueryAPIResponseModel).Reset()
}

// AlibabaIdleIsvUserQueryAPIResponseModel is 服务商ISV闲鱼用户信息查询 成功返回结果
type AlibabaIdleIsvUserQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_user_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleIsvUserQueryTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvUserQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvUserQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvUserQueryAPIResponse)
	},
}

// GetAlibabaIdleIsvUserQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvUserQueryAPIResponse
func GetAlibabaIdleIsvUserQueryAPIResponse() *AlibabaIdleIsvUserQueryAPIResponse {
	return poolAlibabaIdleIsvUserQueryAPIResponse.Get().(*AlibabaIdleIsvUserQueryAPIResponse)
}

// ReleaseAlibabaIdleIsvUserQueryAPIResponse 将 AlibabaIdleIsvUserQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvUserQueryAPIResponse(v *AlibabaIdleIsvUserQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvUserQueryAPIResponse.Put(v)
}
