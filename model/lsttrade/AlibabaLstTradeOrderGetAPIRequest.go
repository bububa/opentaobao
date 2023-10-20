package lsttrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradeordergetAPIRequest 零售通交易订单查询--品牌商视角 API请求
// alibaba.lst.trade.order.get
//
// 根据订单id查询零售通交易订单
type AlibabalsttradeordergetAPIRequest struct {
	model.Params
	// 主订单id
	_mainOrderId int64
	// 子订单id
	_subOrderId int64
}

// NewAlibabalsttradeordergetRequest 初始化AlibabalsttradeordergetAPIRequest对象
func NewAlibabalsttradeordergetRequest() *AlibabalsttradeordergetAPIRequest {
	return &AlibabalsttradeordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradeordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradeordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradeordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMainOrderId is MainOrderId Setter
// 主订单id
func (r *AlibabalsttradeordergetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r AlibabalsttradeordergetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetSubOrderId is SubOrderId Setter
// 子订单id
func (r *AlibabalsttradeordergetAPIRequest) SetSubOrderId(_subOrderId int64) error {
	r._subOrderId = _subOrderId
	r.Set("sub_order_id", _subOrderId)
	return nil
}

// GetSubOrderId SubOrderId Getter
func (r AlibabalsttradeordergetAPIRequest) GetSubOrderId() int64 {
	return r._subOrderId
}
