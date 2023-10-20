package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleConsignmentiiOrderPerformAPIResponse 寄卖V2订单履约 API返回值
// alibaba.idle.consignmentii.order.perform
//
// 寄卖V2订单履约，服务商同步订单信息，驱动交易流转
type AlibabaIdleConsignmentiiOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleConsignmentiiOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentiiOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleConsignmentiiOrderPerformAPIResponseModel).Reset()
}

// AlibabaIdleConsignmentiiOrderPerformAPIResponseModel is 寄卖V2订单履约 成功返回结果
type AlibabaIdleConsignmentiiOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignmentii_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleConsignmentiiOrderPerformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleConsignmentiiOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleConsignmentiiOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleConsignmentiiOrderPerformAPIResponse)
	},
}

// GetAlibabaIdleConsignmentiiOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleConsignmentiiOrderPerformAPIResponse
func GetAlibabaIdleConsignmentiiOrderPerformAPIResponse() *AlibabaIdleConsignmentiiOrderPerformAPIResponse {
	return poolAlibabaIdleConsignmentiiOrderPerformAPIResponse.Get().(*AlibabaIdleConsignmentiiOrderPerformAPIResponse)
}

// ReleaseAlibabaIdleConsignmentiiOrderPerformAPIResponse 将 AlibabaIdleConsignmentiiOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleConsignmentiiOrderPerformAPIResponse(v *AlibabaIdleConsignmentiiOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleConsignmentiiOrderPerformAPIResponse.Put(v)
}
