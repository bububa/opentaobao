package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallchanneltraderefundordergetsAPIRequest 供应商查询退款单 API请求
// tmall.channel.trade.refundorder.gets
//
// 供应商分页查询退款单
type TmallchanneltraderefundordergetsAPIRequest struct {
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

// NewTmallchanneltraderefundordergetsRequest 初始化TmallchanneltraderefundordergetsAPIRequest对象
func NewTmallchanneltraderefundordergetsRequest() *TmallchanneltraderefundordergetsAPIRequest {
	return &TmallchanneltraderefundordergetsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallchanneltraderefundordergetsAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.refundorder.gets"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallchanneltraderefundordergetsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallchanneltraderefundordergetsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainChannelOrderNo is MainChannelOrderNo Setter
// 采购单号
func (r *TmallchanneltraderefundordergetsAPIRequest) SetMainChannelOrderNo(_mainChannelOrderNo string) error {
	r._mainChannelOrderNo = _mainChannelOrderNo
	r.Set("main_channel_order_no", _mainChannelOrderNo)
	return nil
}

// GetMainChannelOrderNo MainChannelOrderNo Getter
func (r TmallchanneltraderefundordergetsAPIRequest) GetMainChannelOrderNo() string {
	return r._mainChannelOrderNo
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TmallchanneltraderefundordergetsAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TmallchanneltraderefundordergetsAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetPageSize is PageSize Setter
// 每页数据条数
func (r *TmallchanneltraderefundordergetsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TmallchanneltraderefundordergetsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNumber is PageNumber Setter
// 页码
func (r *TmallchanneltraderefundordergetsAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TmallchanneltraderefundordergetsAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}
