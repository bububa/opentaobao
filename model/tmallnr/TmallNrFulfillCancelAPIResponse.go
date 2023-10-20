package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillCancelAPIResponse 取消上门揽件 API返回值
// tmall.nr.fulfill.cancel
//
// 新零售门店业务取消上门揽件
type TmallNrFulfillCancelAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillCancelAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillCancelAPIResponseModel).Reset()
}

// TmallNrFulfillCancelAPIResponseModel is 取消上门揽件 成功返回结果
type TmallNrFulfillCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillCancelAPIResponse)
	},
}

// GetTmallNrFulfillCancelAPIResponse 从 sync.Pool 获取 TmallNrFulfillCancelAPIResponse
func GetTmallNrFulfillCancelAPIResponse() *TmallNrFulfillCancelAPIResponse {
	return poolTmallNrFulfillCancelAPIResponse.Get().(*TmallNrFulfillCancelAPIResponse)
}

// ReleaseTmallNrFulfillCancelAPIResponse 将 TmallNrFulfillCancelAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillCancelAPIResponse(v *TmallNrFulfillCancelAPIResponse) {
	v.Reset()
	poolTmallNrFulfillCancelAPIResponse.Put(v)
}
