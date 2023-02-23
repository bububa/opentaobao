package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePayoutGetAPIRequest 获取赔付计时列表数据 API请求
// alibaba.einvoice.payout.get
//
// 获取赔付计时列表数据
type AlibabaEinvoicePayoutGetAPIRequest struct {
	model.Params
	// 当前页码
	_pageNo int64
	// 每页大小，最大50
	_pageSize int64
}

// NewAlibabaEinvoicePayoutGetRequest 初始化AlibabaEinvoicePayoutGetAPIRequest对象
func NewAlibabaEinvoicePayoutGetRequest() *AlibabaEinvoicePayoutGetAPIRequest {
	return &AlibabaEinvoicePayoutGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoicePayoutGetAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.payout.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoicePayoutGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoicePayoutGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 当前页码
func (r *AlibabaEinvoicePayoutGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaEinvoicePayoutGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小，最大50
func (r *AlibabaEinvoicePayoutGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaEinvoicePayoutGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
