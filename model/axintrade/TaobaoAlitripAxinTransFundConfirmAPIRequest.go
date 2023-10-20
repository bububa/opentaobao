package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripaxintransfundconfirmAPIRequest 确认资金单 API请求
// taobao.alitrip.axin.trans.fund.confirm
//
// 通过外部订单号进行资金结算
type TaobaoalitripaxintransfundconfirmAPIRequest struct {
	model.Params
	// 外部订单编号
	_outerOrderId string
}

// NewTaobaoalitripaxintransfundconfirmRequest 初始化TaobaoalitripaxintransfundconfirmAPIRequest对象
func NewTaobaoalitripaxintransfundconfirmRequest() *TaobaoalitripaxintransfundconfirmAPIRequest {
	return &TaobaoalitripaxintransfundconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripaxintransfundconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.axin.trans.fund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripaxintransfundconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripaxintransfundconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterOrderId is OuterOrderId Setter
// 外部订单编号
func (r *TaobaoalitripaxintransfundconfirmAPIRequest) SetOuterOrderId(_outerOrderId string) error {
	r._outerOrderId = _outerOrderId
	r.Set("outer_order_id", _outerOrderId)
	return nil
}

// GetOuterOrderId OuterOrderId Getter
func (r TaobaoalitripaxintransfundconfirmAPIRequest) GetOuterOrderId() string {
	return r._outerOrderId
}
