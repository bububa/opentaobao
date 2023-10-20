package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacsrdonateinvoicequerytoblockchainossAPIRequest 触发odps任务离线查询公益宝贝开票对账明细 API请求
// alibaba.csr.donate.invoice.querytoblockchainoss
//
// 提供给蚂蚁链上公益团队，用于触发odps任务离线查询公益宝贝开票对账明细
type AlibabacsrdonateinvoicequerytoblockchainossAPIRequest struct {
	model.Params
	// 公益宝贝开票对账信息查询入参
	_accountCheckQuery *AccountCheckQuery
}

// NewAlibabacsrdonateinvoicequerytoblockchainossRequest 初始化AlibabacsrdonateinvoicequerytoblockchainossAPIRequest对象
func NewAlibabacsrdonateinvoicequerytoblockchainossRequest() *AlibabacsrdonateinvoicequerytoblockchainossAPIRequest {
	return &AlibabacsrdonateinvoicequerytoblockchainossAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacsrdonateinvoicequerytoblockchainossAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.invoice.querytoblockchainoss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacsrdonateinvoicequerytoblockchainossAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacsrdonateinvoicequerytoblockchainossAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountCheckQuery is AccountCheckQuery Setter
// 公益宝贝开票对账信息查询入参
func (r *AlibabacsrdonateinvoicequerytoblockchainossAPIRequest) SetAccountCheckQuery(_accountCheckQuery *AccountCheckQuery) error {
	r._accountCheckQuery = _accountCheckQuery
	r.Set("account_check_query", _accountCheckQuery)
	return nil
}

// GetAccountCheckQuery AccountCheckQuery Getter
func (r AlibabacsrdonateinvoicequerytoblockchainossAPIRequest) GetAccountCheckQuery() *AccountCheckQuery {
	return r._accountCheckQuery
}
