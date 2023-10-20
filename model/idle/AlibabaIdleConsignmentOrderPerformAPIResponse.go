package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentOrderPerformAPIResponse 帮卖订单履约 API返回值
// alibaba.idle.consignment.order.perform
//
// 帮卖订单履约，回收商同步订单信息，驱动交易流转
type AlibabaIdleConsignmentOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleConsignmentOrderPerformAPIResponseModel).Reset()
}

// AlibabaIdleConsignmentOrderPerformAPIResponseModel is 帮卖订单履约 成功返回结果
type AlibabaIdleConsignmentOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignment_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleConsignmentOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleConsignmentOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentOrderPerformAPIResponse)
	},
}

// GetAlibabaIdleConsignmentOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleConsignmentOrderPerformAPIResponse
func GetAlibabaIdleConsignmentOrderPerformAPIResponse() *AlibabaIdleConsignmentOrderPerformAPIResponse {
	return poolAlibabaIdleConsignmentOrderPerformAPIResponse.Get().(*AlibabaIdleConsignmentOrderPerformAPIResponse)
}

// ReleaseAlibabaIdleConsignmentOrderPerformAPIResponse 将 AlibabaIdleConsignmentOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleConsignmentOrderPerformAPIResponse(v *AlibabaIdleConsignmentOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleConsignmentOrderPerformAPIResponse.Put(v)
}
