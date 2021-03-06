package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeleteprojectAPIResponse 大麦换验平台-第三方对外开放-项目接口deleteProject API返回值
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
type AlibabaDamaiMevOpenDeleteprojectAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteprojectAPIResponseModel
}

// AlibabaDamaiMevOpenDeleteprojectAPIResponseModel is 大麦换验平台-第三方对外开放-项目接口deleteProject 成功返回结果
type AlibabaDamaiMevOpenDeleteprojectAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteproject_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteprojectResult `json:"result,omitempty" xml:"result,omitempty"`
}
