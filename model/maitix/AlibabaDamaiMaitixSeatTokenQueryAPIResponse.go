package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixSeatTokenQueryAPIResponse 分销商选座获取qtoken API返回值
// alibaba.damai.maitix.seat.token.query
//
// 选座分销，分销商查询token
type AlibabaDamaiMaitixSeatTokenQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixSeatTokenQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixSeatTokenQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixSeatTokenQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixSeatTokenQueryAPIResponseModel is 分销商选座获取qtoken 成功返回结果
type AlibabaDamaiMaitixSeatTokenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_seat_token_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixSeatTokenQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixSeatTokenQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixSeatTokenQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixSeatTokenQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixSeatTokenQueryAPIResponse
func GetAlibabaDamaiMaitixSeatTokenQueryAPIResponse() *AlibabaDamaiMaitixSeatTokenQueryAPIResponse {
	return poolAlibabaDamaiMaitixSeatTokenQueryAPIResponse.Get().(*AlibabaDamaiMaitixSeatTokenQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixSeatTokenQueryAPIResponse 将 AlibabaDamaiMaitixSeatTokenQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixSeatTokenQueryAPIResponse(v *AlibabaDamaiMaitixSeatTokenQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixSeatTokenQueryAPIResponse.Put(v)
}
