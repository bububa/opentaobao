package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeRefundorderGetsAPIRequest 供应商查询退款单 API请求
// tmall.channel.trade.refundorder.gets
//
// 供应商分页查询退款单
type TmallChannelTradeRefundorderGetsAPIRequest struct {
	model.Params
	// 采购单号
	_mainChannelOrderNo string
	// 退款单号
	_refundId int64
	// 每页数据条数
	_pageSize int64
	// 页码
	_pageNumber int64
}

// NewTmallChannelTradeRefundorderGetsRequest 初始化TmallChannelTradeRefundorderGetsAPIRequest对象
func NewTmallChannelTradeRefundorderGetsRequest() *TmallChannelTradeRefundorderGetsAPIRequest {
	return &TmallChannelTradeRefundorderGetsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.refundorder.gets"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMainChannelOrderNo is MainChannelOrderNo Setter
// 采购单号
func (r *TmallChannelTradeRefundorderGetsAPIRequest) SetMainChannelOrderNo(_mainChannelOrderNo string) error {
	r._mainChannelOrderNo = _mainChannelOrderNo
	r.Set("main_channel_order_no", _mainChannelOrderNo)
	return nil
}

// GetMainChannelOrderNo MainChannelOrderNo Getter
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetMainChannelOrderNo() string {
	return r._mainChannelOrderNo
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TmallChannelTradeRefundorderGetsAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetPageSize is PageSize Setter
// 每页数据条数
func (r *TmallChannelTradeRefundorderGetsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TmallChannelTradeRefundorderGetsAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TmallChannelTradeRefundorderGetsAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
