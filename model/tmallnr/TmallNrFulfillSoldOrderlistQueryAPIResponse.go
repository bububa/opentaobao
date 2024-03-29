package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillSoldOrderlistQueryAPIResponse 零售商获取品牌商的特定订单列表 API返回值
// tmall.nr.fulfill.sold.orderlist.query
//
// 零售商获取品牌商的特定订单列表，后端服务有零售商和品牌商的绑定关系，存在开关控制；同时后端存在定时送业务等特殊业务的校验，非同城配送业务不能返回，返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位
type TmallNrFulfillSoldOrderlistQueryAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillSoldOrderlistQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillSoldOrderlistQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillSoldOrderlistQueryAPIResponseModel).Reset()
}

// TmallNrFulfillSoldOrderlistQueryAPIResponseModel is 零售商获取品牌商的特定订单列表 成功返回结果
type TmallNrFulfillSoldOrderlistQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_sold_orderlist_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillSoldOrderlistQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillSoldOrderlistQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillSoldOrderlistQueryAPIResponse)
	},
}

// GetTmallNrFulfillSoldOrderlistQueryAPIResponse 从 sync.Pool 获取 TmallNrFulfillSoldOrderlistQueryAPIResponse
func GetTmallNrFulfillSoldOrderlistQueryAPIResponse() *TmallNrFulfillSoldOrderlistQueryAPIResponse {
	return poolTmallNrFulfillSoldOrderlistQueryAPIResponse.Get().(*TmallNrFulfillSoldOrderlistQueryAPIResponse)
}

// ReleaseTmallNrFulfillSoldOrderlistQueryAPIResponse 将 TmallNrFulfillSoldOrderlistQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillSoldOrderlistQueryAPIResponse(v *TmallNrFulfillSoldOrderlistQueryAPIResponse) {
	v.Reset()
	poolTmallNrFulfillSoldOrderlistQueryAPIResponse.Put(v)
}
