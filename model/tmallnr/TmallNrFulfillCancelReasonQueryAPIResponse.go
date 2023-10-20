package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillCancelReasonQueryAPIResponse 查询取消履约的原因列表 API返回值
// tmall.nr.fulfill.cancel.reason.query
//
// 新零售门店业务查询取消上门揽件的原因列表
type TmallNrFulfillCancelReasonQueryAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillCancelReasonQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillCancelReasonQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillCancelReasonQueryAPIResponseModel).Reset()
}

// TmallNrFulfillCancelReasonQueryAPIResponseModel is 查询取消履约的原因列表 成功返回结果
type TmallNrFulfillCancelReasonQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_reason_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillCancelReasonQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillCancelReasonQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillCancelReasonQueryAPIResponse)
	},
}

// GetTmallNrFulfillCancelReasonQueryAPIResponse 从 sync.Pool 获取 TmallNrFulfillCancelReasonQueryAPIResponse
func GetTmallNrFulfillCancelReasonQueryAPIResponse() *TmallNrFulfillCancelReasonQueryAPIResponse {
	return poolTmallNrFulfillCancelReasonQueryAPIResponse.Get().(*TmallNrFulfillCancelReasonQueryAPIResponse)
}

// ReleaseTmallNrFulfillCancelReasonQueryAPIResponse 将 TmallNrFulfillCancelReasonQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillCancelReasonQueryAPIResponse(v *TmallNrFulfillCancelReasonQueryAPIResponse) {
	v.Reset()
	poolTmallNrFulfillCancelReasonQueryAPIResponse.Put(v)
}
