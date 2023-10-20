package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomemianuserbindAPIResponse 主账号入驻 API返回值
// alibaba.alihouse.existinghome.mian.user.bind
//
// 主账号入驻
type AlibabaalihouseexistinghomemianuserbindAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomemianuserbindAPIResponseModel
}

// AlibabaalihouseexistinghomemianuserbindAPIResponseModel is 主账号入驻 成功返回结果
type AlibabaalihouseexistinghomemianuserbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_mian_user_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaalihouseexistinghomemianuserbindResult `json:"result,omitempty" xml:"result,omitempty"`
}
