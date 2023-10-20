package tmallcar

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentAuthCheckAPIResponse 商家鉴权 API返回值
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
type TmallAliautoFulfillmentAuthCheckAPIResponse struct {
	model.CommonResponse
	TmallAliautoFulfillmentAuthCheckAPIResponseModel
}

// TmallAliautoFulfillmentAuthCheckAPIResponseModel is 商家鉴权 成功返回结果
type TmallAliautoFulfillmentAuthCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_fulfillment_auth_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *AliAutoResult `json:"result,omitempty" xml:"result,omitempty"`
}
