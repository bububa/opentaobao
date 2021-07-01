package ascpffo

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressFulfillmentEventAPIResponse
AE履约事件处理 API返回值
aliexpress.fulfillment.event

AE用 履约底层声明发货能力 */
type AliexpressFulfillmentEventAPIResponse struct {
	model.CommonResponse
	AliexpressFulfillmentEventAPIResponseModel
}

// AliexpressFulfillmentEventAPIResponseModel is AE履约事件处理 成功返回结果
type AliexpressFulfillmentEventAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_fulfillment_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AliexpressFulfillmentEventResult `json:"result,omitempty" xml:"result,omitempty"`
}
