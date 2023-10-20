package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderDealrefundAPIResponse 闲鱼无忧购入仓模式服务商退款处理接口 API返回值
// alibaba.idle.isv.order.dealrefund
//
// 闲鱼无忧购业务入仓模式下，用户发起退款后，服务商使用此接口处理退款
type AlibabaIdleIsvOrderDealrefundAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvOrderDealrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderDealrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvOrderDealrefundAPIResponseModel).Reset()
}

// AlibabaIdleIsvOrderDealrefundAPIResponseModel is 闲鱼无忧购入仓模式服务商退款处理接口 成功返回结果
type AlibabaIdleIsvOrderDealrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_dealrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 退款信息
	Module *RefundDto `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderDealrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = nil
}

var poolAlibabaIdleIsvOrderDealrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderDealrefundAPIResponse)
	},
}

// GetAlibabaIdleIsvOrderDealrefundAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvOrderDealrefundAPIResponse
func GetAlibabaIdleIsvOrderDealrefundAPIResponse() *AlibabaIdleIsvOrderDealrefundAPIResponse {
	return poolAlibabaIdleIsvOrderDealrefundAPIResponse.Get().(*AlibabaIdleIsvOrderDealrefundAPIResponse)
}

// ReleaseAlibabaIdleIsvOrderDealrefundAPIResponse 将 AlibabaIdleIsvOrderDealrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvOrderDealrefundAPIResponse(v *AlibabaIdleIsvOrderDealrefundAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvOrderDealrefundAPIResponse.Put(v)
}
