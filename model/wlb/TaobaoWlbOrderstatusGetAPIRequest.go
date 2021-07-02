package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOrderstatusGetAPIRequest 物流宝订单流转状态查询 API请求
// taobao.wlb.orderstatus.get
//
// 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
type TaobaoWlbOrderstatusGetAPIRequest struct {
	model.Params
	// 物流宝订单编码
	_orderCode string
}

// NewTaobaoWlbOrderstatusGetRequest 初始化TaobaoWlbOrderstatusGetAPIRequest对象
func NewTaobaoWlbOrderstatusGetRequest() *TaobaoWlbOrderstatusGetAPIRequest {
	return &TaobaoWlbOrderstatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOrderstatusGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.orderstatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOrderstatusGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderCode is OrderCode Setter
// 物流宝订单编码
func (r *TaobaoWlbOrderstatusGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbOrderstatusGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}
