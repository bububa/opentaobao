package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeOrderGetAPIRequest 通过主采购单号查询采购单 API请求
// tmall.channel.trade.order.get
//
// 通过主采购单号查询采购单
type TmallChannelTradeOrderGetAPIRequest struct {
	model.Params
	// 主采购单ID
	_mainPurchaseOrderNo int64
	// 是否包含子采购单
	_isIncludeSubOrder bool
	// 是否包含主采购单（针对特殊业务）
	_isIncludeMainOrder bool
	// 是否包含物流信息
	_isIncludeLogistics bool
}

// NewTmallChannelTradeOrderGetRequest 初始化TmallChannelTradeOrderGetAPIRequest对象
func NewTmallChannelTradeOrderGetRequest() *TmallChannelTradeOrderGetAPIRequest {
	return &TmallChannelTradeOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallChannelTradeOrderGetAPIRequest) GetApiMethodName() string {
	return "tmall.channel.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallChannelTradeOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMainPurchaseOrderNo is MainPurchaseOrderNo Setter
// 主采购单ID
func (r *TmallChannelTradeOrderGetAPIRequest) SetMainPurchaseOrderNo(_mainPurchaseOrderNo int64) error {
	r._mainPurchaseOrderNo = _mainPurchaseOrderNo
	r.Set("main_purchase_order_no", _mainPurchaseOrderNo)
	return nil
}

// GetMainPurchaseOrderNo MainPurchaseOrderNo Getter
func (r TmallChannelTradeOrderGetAPIRequest) GetMainPurchaseOrderNo() int64 {
	return r._mainPurchaseOrderNo
}

// SetIsIncludeSubOrder is IsIncludeSubOrder Setter
// 是否包含子采购单
func (r *TmallChannelTradeOrderGetAPIRequest) SetIsIncludeSubOrder(_isIncludeSubOrder bool) error {
	r._isIncludeSubOrder = _isIncludeSubOrder
	r.Set("is_include_sub_order", _isIncludeSubOrder)
	return nil
}

// GetIsIncludeSubOrder IsIncludeSubOrder Getter
func (r TmallChannelTradeOrderGetAPIRequest) GetIsIncludeSubOrder() bool {
	return r._isIncludeSubOrder
}

// SetIsIncludeMainOrder is IsIncludeMainOrder Setter
// 是否包含主采购单（针对特殊业务）
func (r *TmallChannelTradeOrderGetAPIRequest) SetIsIncludeMainOrder(_isIncludeMainOrder bool) error {
	r._isIncludeMainOrder = _isIncludeMainOrder
	r.Set("is_include_main_order", _isIncludeMainOrder)
	return nil
}

// GetIsIncludeMainOrder IsIncludeMainOrder Getter
func (r TmallChannelTradeOrderGetAPIRequest) GetIsIncludeMainOrder() bool {
	return r._isIncludeMainOrder
}

// SetIsIncludeLogistics is IsIncludeLogistics Setter
// 是否包含物流信息
func (r *TmallChannelTradeOrderGetAPIRequest) SetIsIncludeLogistics(_isIncludeLogistics bool) error {
	r._isIncludeLogistics = _isIncludeLogistics
	r.Set("is_include_logistics", _isIncludeLogistics)
	return nil
}

// GetIsIncludeLogistics IsIncludeLogistics Getter
func (r TmallChannelTradeOrderGetAPIRequest) GetIsIncludeLogistics() bool {
	return r._isIncludeLogistics
}
