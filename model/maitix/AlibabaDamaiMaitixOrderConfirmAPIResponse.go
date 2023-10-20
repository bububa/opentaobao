package maitix

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderConfirmAPIResponse 大麦-出票 API返回值
// alibaba.damai.maitix.order.confirm
//
// 出票
type AlibabaDamaiMaitixOrderConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOrderConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDamaiMaitixOrderConfirmAPIResponseModel).Reset()
}

// AlibabaDamaiMaitixOrderConfirmAPIResponseModel is 大麦-出票 成功返回结果
type AlibabaDamaiMaitixOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDamaiMaitixOrderConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDamaiMaitixOrderConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOrderConfirmAPIResponse)
	},
}

// GetAlibabaDamaiMaitixOrderConfirmAPIResponse 从 sync.Pool 获取 AlibabaDamaiMaitixOrderConfirmAPIResponse
func GetAlibabaDamaiMaitixOrderConfirmAPIResponse() *AlibabaDamaiMaitixOrderConfirmAPIResponse {
	return poolAlibabaDamaiMaitixOrderConfirmAPIResponse.Get().(*AlibabaDamaiMaitixOrderConfirmAPIResponse)
}

// ReleaseAlibabaDamaiMaitixOrderConfirmAPIResponse 将 AlibabaDamaiMaitixOrderConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDamaiMaitixOrderConfirmAPIResponse(v *AlibabaDamaiMaitixOrderConfirmAPIResponse) {
	v.Reset()
	poolAlibabaDamaiMaitixOrderConfirmAPIResponse.Put(v)
}
