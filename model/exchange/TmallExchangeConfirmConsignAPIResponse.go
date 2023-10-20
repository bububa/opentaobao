package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeConfirmConsignAPIResponse 换货商家确认收货并发货 API返回值
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
type TmallExchangeConfirmConsignAPIResponse struct {
	model.CommonResponse
	TmallExchangeConfirmConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeConfirmConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeConfirmConsignAPIResponseModel).Reset()
}

// TmallExchangeConfirmConsignAPIResponseModel is 换货商家确认收货并发货 成功返回结果
type TmallExchangeConfirmConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_confirm_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeConfirmConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeConfirmConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeConfirmConsignAPIResponse)
	},
}

// GetTmallExchangeConfirmConsignAPIResponse 从 sync.Pool 获取 TmallExchangeConfirmConsignAPIResponse
func GetTmallExchangeConfirmConsignAPIResponse() *TmallExchangeConfirmConsignAPIResponse {
	return poolTmallExchangeConfirmConsignAPIResponse.Get().(*TmallExchangeConfirmConsignAPIResponse)
}

// ReleaseTmallExchangeConfirmConsignAPIResponse 将 TmallExchangeConfirmConsignAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeConfirmConsignAPIResponse(v *TmallExchangeConfirmConsignAPIResponse) {
	v.Reset()
	poolTmallExchangeConfirmConsignAPIResponse.Put(v)
}
