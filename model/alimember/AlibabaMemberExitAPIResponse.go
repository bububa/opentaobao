package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberExitAPIResponse 退会 API返回值
// alibaba.member.exit
//
// 商家会员解绑
type AlibabaMemberExitAPIResponse struct {
	model.CommonResponse
	AlibabaMemberExitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberExitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberExitAPIResponseModel).Reset()
}

// AlibabaMemberExitAPIResponseModel is 退会 成功返回结果
type AlibabaMemberExitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_exit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaMemberExitResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberExitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberExitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberExitAPIResponse)
	},
}

// GetAlibabaMemberExitAPIResponse 从 sync.Pool 获取 AlibabaMemberExitAPIResponse
func GetAlibabaMemberExitAPIResponse() *AlibabaMemberExitAPIResponse {
	return poolAlibabaMemberExitAPIResponse.Get().(*AlibabaMemberExitAPIResponse)
}

// ReleaseAlibabaMemberExitAPIResponse 将 AlibabaMemberExitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberExitAPIResponse(v *AlibabaMemberExitAPIResponse) {
	v.Reset()
	poolAlibabaMemberExitAPIResponse.Put(v)
}
