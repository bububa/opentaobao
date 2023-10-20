package user

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBeneiftDrawAPIResponse 抽奖接口 API返回值
// alibaba.beneift.draw
//
// 抽奖接口
type AlibabaBeneiftDrawAPIResponse struct {
	model.CommonResponse
	AlibabaBeneiftDrawAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBeneiftDrawAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBeneiftDrawAPIResponseModel).Reset()
}

// AlibabaBeneiftDrawAPIResponseModel is 抽奖接口 成功返回结果
type AlibabaBeneiftDrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_beneift_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 权益id
	RightId string `json:"right_id,omitempty" xml:"right_id,omitempty"`
	// 奖品id
	PrizeId string `json:"prize_id,omitempty" xml:"prize_id,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBeneiftDrawAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = ""
	m.RightId = ""
	m.PrizeId = ""
	m.ResultSuccess = false
}

var poolAlibabaBeneiftDrawAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBeneiftDrawAPIResponse)
	},
}

// GetAlibabaBeneiftDrawAPIResponse 从 sync.Pool 获取 AlibabaBeneiftDrawAPIResponse
func GetAlibabaBeneiftDrawAPIResponse() *AlibabaBeneiftDrawAPIResponse {
	return poolAlibabaBeneiftDrawAPIResponse.Get().(*AlibabaBeneiftDrawAPIResponse)
}

// ReleaseAlibabaBeneiftDrawAPIResponse 将 AlibabaBeneiftDrawAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBeneiftDrawAPIResponse(v *AlibabaBeneiftDrawAPIResponse) {
	v.Reset()
	poolAlibabaBeneiftDrawAPIResponse.Put(v)
}
