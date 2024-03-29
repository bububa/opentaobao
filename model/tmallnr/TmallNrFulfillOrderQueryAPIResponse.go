package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillOrderQueryAPIResponse 零售商获取品牌商的单笔订单 API返回值
// tmall.nr.fulfill.order.query
//
// 零售商获取品牌商的单笔订单，后端服务有零售商和品牌商的绑定关系，存在开关控制；返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
type TmallNrFulfillOrderQueryAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillOrderQueryAPIResponseModel).Reset()
}

// TmallNrFulfillOrderQueryAPIResponseModel is 零售商获取品牌商的单笔订单 成功返回结果
type TmallNrFulfillOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillOrderQueryAPIResponse)
	},
}

// GetTmallNrFulfillOrderQueryAPIResponse 从 sync.Pool 获取 TmallNrFulfillOrderQueryAPIResponse
func GetTmallNrFulfillOrderQueryAPIResponse() *TmallNrFulfillOrderQueryAPIResponse {
	return poolTmallNrFulfillOrderQueryAPIResponse.Get().(*TmallNrFulfillOrderQueryAPIResponse)
}

// ReleaseTmallNrFulfillOrderQueryAPIResponse 将 TmallNrFulfillOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillOrderQueryAPIResponse(v *TmallNrFulfillOrderQueryAPIResponse) {
	v.Reset()
	poolTmallNrFulfillOrderQueryAPIResponse.Put(v)
}
