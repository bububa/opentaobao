package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushitemAPIResponse 大麦换验平台-第三方对外开放-票品接口pushItem API返回值
// alibaba.damai.mev.open.pushitem
//
// 开放接口 推送票品
type AlibabadamaimevopenpushitemAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushitemAPIResponseModel
}

// AlibabadamaimevopenpushitemAPIResponseModel is 大麦换验平台-第三方对外开放-票品接口pushItem 成功返回结果
type AlibabadamaimevopenpushitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushitemResult `json:"result,omitempty" xml:"result,omitempty"`
}
