package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomepaymentmethodsyncAPIResponse 付款方式上翻 API返回值
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
type AlibabaalihousenewhomepaymentmethodsyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomepaymentmethodsyncAPIResponseModel
}

// AlibabaalihousenewhomepaymentmethodsyncAPIResponseModel is 付款方式上翻 成功返回结果
type AlibabaalihousenewhomepaymentmethodsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_payment_method_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AlibabaalihousenewhomepaymentmethodsyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
