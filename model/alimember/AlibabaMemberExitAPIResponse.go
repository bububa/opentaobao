package alimember

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberexitAPIResponse 退会 API返回值
// alibaba.member.exit
//
// 商家会员解绑
type AlibabamemberexitAPIResponse struct {
	model.CommonResponse
	AlibabamemberexitAPIResponseModel
}

// AlibabamemberexitAPIResponseModel is 退会 成功返回结果
type AlibabamemberexitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_exit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabamemberexitResult `json:"result,omitempty" xml:"result,omitempty"`
}
