package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSoldOrderlistQueryJstAPIResponse 已入塔商家查询订单列表 API返回值
// tmall.nr.sold.orderlist.query.jst
//
// 该服务用于已入聚石塔的商家，获取订单列表信息；
type TmallNrSoldOrderlistQueryJstAPIResponse struct {
	model.CommonResponse
	TmallNrSoldOrderlistQueryJstAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrSoldOrderlistQueryJstAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrSoldOrderlistQueryJstAPIResponseModel).Reset()
}

// TmallNrSoldOrderlistQueryJstAPIResponseModel is 已入塔商家查询订单列表 成功返回结果
type TmallNrSoldOrderlistQueryJstAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_sold_orderlist_query_jst_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrSoldOrderlistQueryJstAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrSoldOrderlistQueryJstAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrSoldOrderlistQueryJstAPIResponse)
	},
}

// GetTmallNrSoldOrderlistQueryJstAPIResponse 从 sync.Pool 获取 TmallNrSoldOrderlistQueryJstAPIResponse
func GetTmallNrSoldOrderlistQueryJstAPIResponse() *TmallNrSoldOrderlistQueryJstAPIResponse {
	return poolTmallNrSoldOrderlistQueryJstAPIResponse.Get().(*TmallNrSoldOrderlistQueryJstAPIResponse)
}

// ReleaseTmallNrSoldOrderlistQueryJstAPIResponse 将 TmallNrSoldOrderlistQueryJstAPIResponse 保存到 sync.Pool
func ReleaseTmallNrSoldOrderlistQueryJstAPIResponse(v *TmallNrSoldOrderlistQueryJstAPIResponse) {
	v.Reset()
	poolTmallNrSoldOrderlistQueryJstAPIResponse.Put(v)
}
