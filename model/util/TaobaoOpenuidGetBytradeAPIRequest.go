package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenuidgetbytradeAPIRequest 通过订单获取对应买家的openUID API请求
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
type TaobaoopenuidgetbytradeAPIRequest struct {
	model.Params
	// 订单ID
	_tid int64
}

// NewTaobaoopenuidgetbytradeRequest 初始化TaobaoopenuidgetbytradeAPIRequest对象
func NewTaobaoopenuidgetbytradeRequest() *TaobaoopenuidgetbytradeAPIRequest {
	return &TaobaoopenuidgetbytradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenuidgetbytradeAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get.bytrade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenuidgetbytradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenuidgetbytradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTid is Tid Setter
// 订单ID
func (r *TaobaoopenuidgetbytradeAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenuidgetbytradeAPIRequest) GetTid() int64 {
	return r._tid
}
