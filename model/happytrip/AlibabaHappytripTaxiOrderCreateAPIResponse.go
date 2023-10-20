package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderCreateAPIResponse 用户叫车 API返回值
// alibaba.happytrip.taxi.order.create
//
// 用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
type AlibabaHappytripTaxiOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiOrderCreateAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiOrderCreateAPIResponseModel is 用户叫车 成功返回结果
type AlibabaHappytripTaxiOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 订单创建结果
	Data *OrderCreateResult `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errmsg = ""
	m.Errno = 0
	m.Data = nil
}

var poolAlibabaHappytripTaxiOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiOrderCreateAPIResponse)
	},
}

// GetAlibabaHappytripTaxiOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiOrderCreateAPIResponse
func GetAlibabaHappytripTaxiOrderCreateAPIResponse() *AlibabaHappytripTaxiOrderCreateAPIResponse {
	return poolAlibabaHappytripTaxiOrderCreateAPIResponse.Get().(*AlibabaHappytripTaxiOrderCreateAPIResponse)
}

// ReleaseAlibabaHappytripTaxiOrderCreateAPIResponse 将 AlibabaHappytripTaxiOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderCreateAPIResponse(v *AlibabaHappytripTaxiOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderCreateAPIResponse.Put(v)
}
