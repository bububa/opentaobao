package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushprojectAPIResponse 大麦换验平台-第三方对外开放-项目接口pushProject API返回值
// alibaba.damai.mev.open.pushproject
//
// pushProject
type AlibabadamaimevopenpushprojectAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopenpushprojectAPIResponseModel
}

// AlibabadamaimevopenpushprojectAPIResponseModel is 大麦换验平台-第三方对外开放-项目接口pushProject 成功返回结果
type AlibabadamaimevopenpushprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_pushproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopenpushprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}
