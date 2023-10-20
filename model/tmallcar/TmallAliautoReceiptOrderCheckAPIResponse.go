package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoReceiptOrderCheckAPIResponse 查看工单查询订单是否已付款 API返回值
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
type TmallAliautoReceiptOrderCheckAPIResponse struct {
	model.CommonResponse
	TmallAliautoReceiptOrderCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoReceiptOrderCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoReceiptOrderCheckAPIResponseModel).Reset()
}

// TmallAliautoReceiptOrderCheckAPIResponseModel is 查看工单查询订单是否已付款 成功返回结果
type TmallAliautoReceiptOrderCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_receipt_order_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *CheckReceiptOrderIsPaid4IsvDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoReceiptOrderCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTmallAliautoReceiptOrderCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoReceiptOrderCheckAPIResponse)
	},
}

// GetTmallAliautoReceiptOrderCheckAPIResponse 从 sync.Pool 获取 TmallAliautoReceiptOrderCheckAPIResponse
func GetTmallAliautoReceiptOrderCheckAPIResponse() *TmallAliautoReceiptOrderCheckAPIResponse {
	return poolTmallAliautoReceiptOrderCheckAPIResponse.Get().(*TmallAliautoReceiptOrderCheckAPIResponse)
}

// ReleaseTmallAliautoReceiptOrderCheckAPIResponse 将 TmallAliautoReceiptOrderCheckAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoReceiptOrderCheckAPIResponse(v *TmallAliautoReceiptOrderCheckAPIResponse) {
	v.Reset()
	poolTmallAliautoReceiptOrderCheckAPIResponse.Put(v)
}
