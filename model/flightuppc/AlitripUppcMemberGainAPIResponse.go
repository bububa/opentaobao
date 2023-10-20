package flightuppc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripUppcMemberGainAPIResponse 航司权益数据回流 API返回值
// alitrip.uppc.member.gain
//
// 航司权益数据回流
type AlitripUppcMemberGainAPIResponse struct {
	model.CommonResponse
	AlitripUppcMemberGainAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripUppcMemberGainAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripUppcMemberGainAPIResponseModel).Reset()
}

// AlitripUppcMemberGainAPIResponseModel is 航司权益数据回流 成功返回结果
type AlitripUppcMemberGainAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_uppc_member_gain_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripUppcMemberGainAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripUppcMemberGainAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripUppcMemberGainAPIResponse)
	},
}

// GetAlitripUppcMemberGainAPIResponse 从 sync.Pool 获取 AlitripUppcMemberGainAPIResponse
func GetAlitripUppcMemberGainAPIResponse() *AlitripUppcMemberGainAPIResponse {
	return poolAlitripUppcMemberGainAPIResponse.Get().(*AlitripUppcMemberGainAPIResponse)
}

// ReleaseAlitripUppcMemberGainAPIResponse 将 AlitripUppcMemberGainAPIResponse 保存到 sync.Pool
func ReleaseAlitripUppcMemberGainAPIResponse(v *AlitripUppcMemberGainAPIResponse) {
	v.Reset()
	poolAlitripUppcMemberGainAPIResponse.Put(v)
}
