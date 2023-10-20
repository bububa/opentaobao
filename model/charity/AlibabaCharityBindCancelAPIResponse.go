package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitybindcancelAPIResponse 取消用户绑定 API返回值
// alibaba.charity.bind.cancel
//
// 取消用户绑定
type AlibabacharitybindcancelAPIResponse struct {
	model.CommonResponse
	AlibabacharitybindcancelAPIResponseModel
}

// AlibabacharitybindcancelAPIResponseModel is 取消用户绑定 成功返回结果
type AlibabacharitybindcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_bind_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *ThreehoursResult `json:"result,omitempty" xml:"result,omitempty"`
}
