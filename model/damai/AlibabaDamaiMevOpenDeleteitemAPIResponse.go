package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeleteitemAPIResponse 大麦换验平台-第三方对外开放-票品接口deleteItem API返回值
// alibaba.damai.mev.open.deleteitem
//
// deleteItem
type AlibabadamaimevopendeleteitemAPIResponse struct {
	model.CommonResponse
	AlibabadamaimevopendeleteitemAPIResponseModel
}

// AlibabadamaimevopendeleteitemAPIResponseModel is 大麦换验平台-第三方对外开放-票品接口deleteItem 成功返回结果
type AlibabadamaimevopendeleteitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabadamaimevopendeleteitemResult `json:"result,omitempty" xml:"result,omitempty"`
}
