package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.finance.bill.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderFinanceBillQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
