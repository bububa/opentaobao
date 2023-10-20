package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeletefloorAPIResponse 大麦换验平台-第三方对外开放-楼层接口deleteFloor API返回值
// alibaba.damai.mev.open.deletefloor
//
// deleteFloor
type AlibabadamaimevopendeletefloorAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopendeletefloorAPIResponseModel
}

// AlibabadamaimevopendeletefloorAPIResponseModel is 大麦换验平台-第三方对外开放-楼层接口deleteFloor 成功返回结果
type AlibabadamaimevopendeletefloorAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deletefloor_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopendeletefloorResult `json:"result,omitempty" xml:"result,omitempty"`
}
