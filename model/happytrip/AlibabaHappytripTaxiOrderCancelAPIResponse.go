package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderCancelAPIResponse 取消叫车 API返回值
// alibaba.happytrip.taxi.order.cancel
//
// 取消叫车订单,行程中的订单不能取消
type AlibabaHappytripTaxiOrderCancelAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderCancelAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderCancelAPIResponseModel is 取消叫车 成功返回结果
type AlibabaHappytripTaxiOrderCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 订单取消结果
	Data *OrderCancelResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiOrderCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderCancelAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderCancelAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderCancelAPIResponse
func GetAlibabaHappytripTaxiOrderCancelAPIResponse() *AlibabaHappytripTaxiOrderCancelAPIResponse {
	return poolAlibabaHappytripTaxiOrderCancelAPIResponse.Get().(*AlibabaHappytripTaxiOrderCancelAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderCancelAPIResponse 将 AlibabaHappytripTaxiOrderCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderCancelAPIResponse(v *AlibabaHappytripTaxiOrderCancelAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderCancelAPIResponse.Put(v)
}
