package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderCancelAPIResponse 大麦-库存释放 API返回值
// alibaba.damai.maitix.order.cancel
//
// 库存释放
type AlibabaDamaiMaitixOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOrderCancelAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOrderCancelAPIResponseModel is 大麦-库存释放 成功返回结果
type AlibabaDamaiMaitixOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOrderCancelAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOrderCancelAPIResponse
func GetAlibabaDamaiMaitixOrderCancelAPIResponse() *AlibabaDamaiMaitixOrderCancelAPIResponse {
	return poolAlibabaDamaiMaitixOrderCancelAPIResponse.Get().(*AlibabaDamaiMaitixOrderCancelAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOrderCancelAPIResponse 将 AlibabaDamaiMaitixOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderCancelAPIResponse(v *AlibabaDamaiMaitixOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderCancelAPIResponse.Put(v)
}
