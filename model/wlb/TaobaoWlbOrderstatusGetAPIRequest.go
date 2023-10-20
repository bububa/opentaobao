package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderstatusgetAPIRequest 物流宝订单流转状态查询 API请求
// taobao.wlb.orderstatus.get
//
// 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaowlborderstatusgetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_orderCode string
}

// NewTaobaowlborderstatusgetRequest 初始化TaobaowlborderstatusgetAPIRequest对象
func NewTaobaowlborderstatusgetRequest() *TaobaowlborderstatusgetAPIRequest {
	return &TaobaowlborderstatusgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlborderstatusgetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.orderstatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlborderstatusgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlborderstatusgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 物流宝订单编码
func (r *TaobaowlborderstatusgetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlborderstatusgetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
