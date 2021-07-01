package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalItemUnbindAPIResponse
ISV解绑商品 API返回值
alibaba.alihealth.dental.item.unbind

ISV解绑商品 */
type AlibabaAlihealthDentalItemUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalItemUnbindAPIResponseModel
}

// AlibabaAlihealthDentalItemUnbindAPIResponseModel is ISV解绑商品 成功返回结果
type AlibabaAlihealthDentalItemUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_item_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalItemUnbindMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
