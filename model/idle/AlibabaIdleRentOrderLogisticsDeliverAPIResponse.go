package idle

import (
	"encoding/xml"

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

// AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel is 创建揽收物流 成功返回结果
type AlibabaIdleRentOrderLogisticsDeliverAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_logistics_deliver_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentOrderLogisticsDeliverTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
