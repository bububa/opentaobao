package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometalentbindstoreAPIResponse 达人号门店关系绑定 API返回值
// alibaba.alihouse.newhome.talent.bind.store
//
// 达人号门店关系绑定
type AlibabaalihousenewhometalentbindstoreAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhometalentbindstoreAPIResponseModel
}

// AlibabaalihousenewhometalentbindstoreAPIResponseModel is 达人号门店关系绑定 成功返回结果
type AlibabaalihousenewhometalentbindstoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_talent_bind_store_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定关系结果列表
	Data []StoreTalentBindingResponseDto `json:"data,omitempty" xml:"data>store_talent_binding_response_dto,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// aa
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// aaa
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}
