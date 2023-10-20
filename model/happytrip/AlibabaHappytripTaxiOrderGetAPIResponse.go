package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderGetAPIResponse 订单详情 API返回值
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
type AlibabaHappytripTaxiOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderGetAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderGetAPIResponseModel is 订单详情 成功返回结果
type AlibabaHappytripTaxiOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 订单获取结果
	Data *OrderGetResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderGetAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderGetAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderGetAPIResponse
func GetAlibabaHappytripTaxiOrderGetAPIResponse() *AlibabaHappytripTaxiOrderGetAPIResponse {
	return poolAlibabaHappytripTaxiOrderGetAPIResponse.Get().(*AlibabaHappytripTaxiOrderGetAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderGetAPIResponse 将 AlibabaHappytripTaxiOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderGetAPIResponse(v *AlibabaHappytripTaxiOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderGetAPIResponse.Put(v)
}
