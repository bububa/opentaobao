package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderFinanceBillQueryAPIRequest 资金合规商家账单 API请求
// alibaba.wdk.order.finance.bill.query
//
// 拉取资金合规商家账单
type AlibabaWdkOrderFinanceBillQueryAPIRequest struct {
	model.Params
	// 入参
	_billQueryRequest *WdkOpenOrderFinanceBillQueryRequest
}

// NewAlibabaWdkOrderFinanceBillQueryRequest 初始化AlibabaWdkOrderFinanceBillQueryAPIRequest对象
func NewAlibabaWdkOrderFinanceBillQueryRequest() *AlibabaWdkOrderFinanceBillQueryAPIRequest {
	return &AlibabaWdkOrderFinanceBillQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderFinanceBillQueryAPIRequest) Reset() {
	r._billQueryRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.finance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillQueryRequest is BillQueryRequest Setter
// 入参
func (r *AlibabaWdkOrderFinanceBillQueryAPIRequest) SetBillQueryRequest(_billQueryRequest *WdkOpenOrderFinanceBillQueryRequest) error {
	r._billQueryRequest = _billQueryRequest
	r.Set("bill_query_request", _billQueryRequest)
	return nil
}

// GetBillQueryRequest BillQueryRequest Getter
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetBillQueryRequest() *WdkOpenOrderFinanceBillQueryRequest {
	return r._billQueryRequest
}

var poolAlibabaWdkOrderFinanceBillQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderFinanceBillQueryRequest()
	},
}

// GetAlibabaWdkOrderFinanceBillQueryRequest 从 sync.Pool 获取 AlibabaWdkOrderFinanceBillQueryAPIRequest
func GetAlibabaWdkOrderFinanceBillQueryAPIRequest() *AlibabaWdkOrderFinanceBillQueryAPIRequest {
	return poolAlibabaWdkOrderFinanceBillQueryAPIRequest.Get().(*AlibabaWdkOrderFinanceBillQueryAPIRequest)
}

// ReleaseAlibabaWdkOrderFinanceBillQueryAPIRequest 将 AlibabaWdkOrderFinanceBillQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderFinanceBillQueryAPIRequest(v *AlibabaWdkOrderFinanceBillQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderFinanceBillQueryAPIRequest.Put(v)
}
