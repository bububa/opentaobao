package damai

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenDeleteitemAPIResponse
大麦换验平台-第三方对外开放-票品接口deleteItem API返回值
alibaba.damai.mev.open.deleteitem

deleteItem */
type AlibabaDamaiMevOpenDeleteitemAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMevOpenDeleteitemAPIResponseModel
}

// AlibabaDamaiMevOpenDeleteitemAPIResponseModel is 大麦换验平台-第三方对外开放-票品接口deleteItem 成功返回结果
type AlibabaDamaiMevOpenDeleteitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_mev_open_deleteitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaDamaiMevOpenDeleteitemResult `json:"result,omitempty" xml:"result,omitempty"`
}
