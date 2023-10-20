package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpresspaymentexchangegetAPIResponse getExchange API返回值
// aliexpress.payment.exchange.get
//
// 提供国际汇率服务
type AliexpresspaymentexchangegetAPIResponse struct {
	model.CommonResponse
	AliexpresspaymentexchangegetAPIResponseModel
}

// AliexpresspaymentexchangegetAPIResponseModel is getExchange 成功返回结果
type AliexpresspaymentexchangegetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_payment_exchange_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpresspaymentexchangegetResult `json:"result,omitempty" xml:"result,omitempty"`
}
