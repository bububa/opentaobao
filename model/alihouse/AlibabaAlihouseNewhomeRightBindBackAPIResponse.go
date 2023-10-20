package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomerightbindbackAPIResponse 权限回流 API返回值
// alibaba.alihouse.newhome.right.bind.back
//
// 权限回流
type AlibabaalihousenewhomerightbindbackAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomerightbindbackAPIResponseModel
}

// AlibabaalihousenewhomerightbindbackAPIResponseModel is 权限回流 成功返回结果
type AlibabaalihousenewhomerightbindbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_right_bind_back_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlibabaalihousenewhomerightbindbackResult `json:"result,omitempty" xml:"result,omitempty"`
}
