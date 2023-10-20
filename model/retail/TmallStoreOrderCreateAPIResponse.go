package retail

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallStoreOrderCreateAPIResponse 门店订单创建api API返回值
// tmall.store.order.create
//
// 门店订单创建api
type TmallStoreOrderCreateAPIResponse struct {
	model.CommonResponse
	TmallStoreOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallStoreOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallStoreOrderCreateAPIResponseModel).Reset()
}

// TmallStoreOrderCreateAPIResponseModel is 门店订单创建api 成功返回结果
type TmallStoreOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_store_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// detailResults
	DetailResults []Detailresults `json:"detail_results,omitempty" xml:"detail_results>detailresults,omitempty"`
}

// Reset 清空结构体
func (m *TmallStoreOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DetailResults = m.DetailResults[:0]
}

var poolTmallStoreOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallStoreOrderCreateAPIResponse)
	},
}

// GetTmallStoreOrderCreateAPIResponse 从 sync.Pool 获取 TmallStoreOrderCreateAPIResponse
func GetTmallStoreOrderCreateAPIResponse() *TmallStoreOrderCreateAPIResponse {
	return poolTmallStoreOrderCreateAPIResponse.Get().(*TmallStoreOrderCreateAPIResponse)
}

// ReleaseTmallStoreOrderCreateAPIResponse 将 TmallStoreOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallStoreOrderCreateAPIResponse(v *TmallStoreOrderCreateAPIResponse) {
	v.Reset()
	poolTmallStoreOrderCreateAPIResponse.Put(v)
}
