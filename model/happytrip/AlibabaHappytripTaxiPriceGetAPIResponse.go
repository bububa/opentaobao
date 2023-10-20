package happytrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiPriceGetAPIResponse 获取价格预估信息 API返回值
// alibaba.happytrip.taxi.price.get
//
// 打车价格预估
type AlibabaHappytripTaxiPriceGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiPriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiPriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHappytripTaxiPriceGetAPIResponseModel).Reset()
}

// AlibabaHappytripTaxiPriceGetAPIResponseModel is 获取价格预估信息 成功返回结果
type AlibabaHappytripTaxiPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 价格预估模型
	Data []PriceModel `json:"data,omitempty" xml:"data>price_model,omitempty"`
	// 错误消息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误代码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHappytripTaxiPriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.Errmsg = ""
	m.Errno = 0
}

var poolAlibabaHappytripTaxiPriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHappytripTaxiPriceGetAPIResponse)
	},
}

// GetAlibabaHappytripTaxiPriceGetAPIResponse 从 sync.Pool 获取 AlibabaHappytripTaxiPriceGetAPIResponse
func GetAlibabaHappytripTaxiPriceGetAPIResponse() *AlibabaHappytripTaxiPriceGetAPIResponse {
	return poolAlibabaHappytripTaxiPriceGetAPIResponse.Get().(*AlibabaHappytripTaxiPriceGetAPIResponse)
}

// ReleaseAlibabaHappytripTaxiPriceGetAPIResponse 将 AlibabaHappytripTaxiPriceGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHappytripTaxiPriceGetAPIResponse(v *AlibabaHappytripTaxiPriceGetAPIResponse) {
	v.Reset()
	poolAlibabaHappytripTaxiPriceGetAPIResponse.Put(v)
}
