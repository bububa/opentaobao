package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeletestandAPIResponse
大麦换验平台-第三方对外开放-看台接口deleteStand API返回值
alibaba.damai.mev.open.deletestand

deleteStand */
type AlibabaDamaiMevOpenDeletestandAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeletestandAPIResponseModel
}

// AlibabaDamaiMevOpenDeletestandAPIResponseModel is 大麦换验平台-第三方对外开放-看台接口deleteStand 成功返回结果
type AlibabaDamaiMevOpenDeletestandAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deletestand_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeletestandResult `json:"result,omitempty" xml:"result,omitempty"`
}
