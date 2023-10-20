package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixSeatInfoQueryAPIResponse 分销商查询座位信息 API返回值
// alibaba.damai.maitix.seat.info.query
//
// 分销查询座位文案信息
type AlibabaDamaiMaitixSeatInfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixSeatInfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel is 分销商查询座位信息 成功返回结果
type AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_seat_info_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixSeatInfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixSeatInfoQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixSeatInfoQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixSeatInfoQueryAPIResponse
func GetAlibabaDamaiMaitixSeatInfoQueryAPIResponse() *AlibabaDamaiMaitixSeatInfoQueryAPIResponse {
	return poolAlibabaDamaiMaitixSeatInfoQueryAPIResponse.Get().(*AlibabaDamaiMaitixSeatInfoQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixSeatInfoQueryAPIResponse 将 AlibabaDamaiMaitixSeatInfoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixSeatInfoQueryAPIResponse(v *AlibabaDamaiMaitixSeatInfoQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixSeatInfoQueryAPIResponse.Put(v)
}
