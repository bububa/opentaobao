package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderLogisticsDeliverAPIResponse 创建揽收物流 API返回值
// alibaba.idle.rent.order.logistics.deliver
//
// 创建揽收物流
// 商家去物流公司创建物流订单
type AlibabaIdleRentOrderLogisticsDeliverAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderLogisticsDeliverAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel).Reset()
}

// AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel is 创建揽收物流 成功返回结果
type AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_logistics_deliver_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderLogisticsDeliverTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentOrderLogisticsDeliverAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderLogisticsDeliverAPIResponse)
	},
}

// GetAlibabaIdleRentOrderLogisticsDeliverAPIResponse 从 sync.Pool 获取 AlibabaIdleRentOrderLogisticsDeliverAPIResponse
func GetAlibabaIdleRentOrderLogisticsDeliverAPIResponse() *AlibabaIdleRentOrderLogisticsDeliverAPIResponse {
	return poolAlibabaIdleRentOrderLogisticsDeliverAPIResponse.Get().(*AlibabaIdleRentOrderLogisticsDeliverAPIResponse)
}

// ReleaseAlibabaIdleRentOrderLogisticsDeliverAPIResponse 将 AlibabaIdleRentOrderLogisticsDeliverAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentOrderLogisticsDeliverAPIResponse(v *AlibabaIdleRentOrderLogisticsDeliverAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentOrderLogisticsDeliverAPIResponse.Put(v)
}
