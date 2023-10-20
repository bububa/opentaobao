package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrOrderQueryJstAPIResponse 获取同城配送业务单笔订单 API返回值
// tmall.nr.order.query.jst
//
// 同城配送业务获取单笔订单
type TmallNrOrderQueryJstAPIResponse struct {
	model.CommonResponse
	TmallNrOrderQueryJstAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrOrderQueryJstAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrOrderQueryJstAPIResponseModel).Reset()
}

// TmallNrOrderQueryJstAPIResponseModel is 获取同城配送业务单笔订单 成功返回结果
type TmallNrOrderQueryJstAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_order_query_jst_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrOrderQueryJstAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrOrderQueryJstAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrOrderQueryJstAPIResponse)
	},
}

// GetTmallNrOrderQueryJstAPIResponse 从 sync.Pool 获取 TmallNrOrderQueryJstAPIResponse
func GetTmallNrOrderQueryJstAPIResponse() *TmallNrOrderQueryJstAPIResponse {
	return poolTmallNrOrderQueryJstAPIResponse.Get().(*TmallNrOrderQueryJstAPIResponse)
}

// ReleaseTmallNrOrderQueryJstAPIResponse 将 TmallNrOrderQueryJstAPIResponse 保存到 sync.Pool
func ReleaseTmallNrOrderQueryJstAPIResponse(v *TmallNrOrderQueryJstAPIResponse) {
	v.Reset()
	poolTmallNrOrderQueryJstAPIResponse.Put(v)
}
