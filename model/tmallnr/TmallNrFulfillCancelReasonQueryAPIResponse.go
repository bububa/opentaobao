package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallnrfulfillcancelreasonqueryAPIResponse 查询取消履约的原因列表 API返回值
// tmall.nr.fulfill.cancel.reason.query
//
// 新零售门店业务查询取消上门揽件的原因列表
type TmallnrfulfillcancelreasonqueryAPIResponse struct {
	model.CommonResponse
	TmallnrfulfillcancelreasonqueryAPIResponseModel
}

// TmallnrfulfillcancelreasonqueryAPIResponseModel is 查询取消履约的原因列表 成功返回结果
type TmallnrfulfillcancelreasonqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_reason_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}
