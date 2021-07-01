package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteperformAPIResponse
大麦换验平台-第三方对外开放-场次接口deletePerform API返回值
alibaba.damai.mev.open.deleteperform

deletePerform */
type AlibabaDamaiMevOpenDeleteperformAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteperformAPIResponseModel
}

// AlibabaDamaiMevOpenDeleteperformAPIResponseModel is 大麦换验平台-第三方对外开放-场次接口deletePerform 成功返回结果
type AlibabaDamaiMevOpenDeleteperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteperform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteperformResult `json:"result,omitempty" xml:"result,omitempty"`
}
