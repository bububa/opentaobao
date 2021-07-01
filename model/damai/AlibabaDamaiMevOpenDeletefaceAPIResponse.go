package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletefaceAPIResponse
大麦换验平台-第三方对外开放-票面接口deleteFace API返回值
alibaba.damai.mev.open.deleteface

deleteFace */
type AlibabaDamaiMevOpenDeletefaceAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletefaceAPIResponseModel
}

// AlibabaDamaiMevOpenDeletefaceAPIResponseModel is 大麦换验平台-第三方对外开放-票面接口deleteFace 成功返回结果
type AlibabaDamaiMevOpenDeletefaceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletefaceResult `json:"result,omitempty" xml:"result,omitempty"`
}
