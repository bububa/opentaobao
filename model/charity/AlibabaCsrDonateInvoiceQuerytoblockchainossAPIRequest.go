package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest 触发odps任务离线查询公益宝贝开票对账明细 API请求
// alibaba.csr.donate.invoice.querytoblockchainoss
//
// 提供给蚂蚁链上公益团队，用于触发odps任务离线查询公益宝贝开票对账明细
type AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest struct {
	model.Params
	// 公益宝贝开票对账信息查询入参
	_accountCheckQuery *AccountCheckQuery
}

// NewAlibabaCsrDonateInvoiceQuerytoblockchainossRequest 初始化AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest对象
func NewAlibabaCsrDonateInvoiceQuerytoblockchainossRequest() *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest {
	return &AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.invoice.querytoblockchainoss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAccountCheckQuery is AccountCheckQuery Setter
// 公益宝贝开票对账信息查询入参
func (r *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) SetAccountCheckQuery(_accountCheckQuery *AccountCheckQuery) error {
	r._accountCheckQuery = _accountCheckQuery
	r.Set("account_check_query", _accountCheckQuery)
	return nil
}

// GetAccountCheckQuery AccountCheckQuery Getter
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetAccountCheckQuery() *AccountCheckQuery {
	return r._accountCheckQuery
}
