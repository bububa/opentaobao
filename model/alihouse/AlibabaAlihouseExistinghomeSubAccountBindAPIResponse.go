package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomesubaccountbindAPIResponse 子账号入驻 API返回值
// alibaba.alihouse.existinghome.sub.account.bind
//
// 子账号入驻
type AlibabaalihouseexistinghomesubaccountbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomesubaccountbindAPIResponseModel
}

// AlibabaalihouseexistinghomesubaccountbindAPIResponseModel is 子账号入驻 成功返回结果
type AlibabaalihouseexistinghomesubaccountbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_sub_account_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihouseexistinghomesubaccountbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
