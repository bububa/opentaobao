package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderQueryAPIResponse 大麦-查询分销单 API返回值
// alibaba.damai.maitix.order.query
//
// 查询分销单
type AlibabaDamaiMaitixOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOrderQueryAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOrderQueryAPIResponseModel is 大麦-查询分销单 成功返回结果
type AlibabaDamaiMaitixOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOrderQueryAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOrderQueryAPIResponse
func GetAlibabaDamaiMaitixOrderQueryAPIResponse() *AlibabaDamaiMaitixOrderQueryAPIResponse {
	return poolAlibabaDamaiMaitixOrderQueryAPIResponse.Get().(*AlibabaDamaiMaitixOrderQueryAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOrderQueryAPIResponse 将 AlibabaDamaiMaitixOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderQueryAPIResponse(v *AlibabaDamaiMaitixOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderQueryAPIResponse.Put(v)
}
