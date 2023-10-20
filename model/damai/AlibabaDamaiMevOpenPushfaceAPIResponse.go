package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushfaceAPIResponse 大麦换验平台-第三方对外开放-票面接口pushFace API返回值
// alibaba.damai.mev.open.pushface
//
// pushFace
type AlibabadamaimevopenpushfaceAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushfaceAPIResponseModel
}

// AlibabadamaimevopenpushfaceAPIResponseModel is 大麦换验平台-第三方对外开放-票面接口pushFace 成功返回结果
type AlibabadamaimevopenpushfaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushfaceResult `json:"result,omitempty" xml:"result,omitempty"`
}
