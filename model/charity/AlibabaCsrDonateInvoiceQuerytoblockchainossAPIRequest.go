package charity

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) Reset() {
	r._accountCheckQuery = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetApiMethodName() string {
	return "alibaba.csr.donate.invoice.querytoblockchainoss"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCsrDonateInvoiceQuerytoblockchainossRequest()
	},
}

// GetAlibabaCsrDonateInvoiceQuerytoblockchainossRequest 从 sync.Pool 获取 AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest
func GetAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest() *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest {
	return poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest.Get().(*AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest)
}

// ReleaseAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest 将 AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest 放入 sync.Pool
func ReleaseAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest(v *AlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest) {
	v.Reset()
	poolAlibabaCsrDonateInvoiceQuerytoblockchainossAPIRequest.Put(v)
}
