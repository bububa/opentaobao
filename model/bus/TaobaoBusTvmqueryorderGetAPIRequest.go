package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmqueryordergetAPIRequest 线下自助机查询订单信息 API请求
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
type TaobaobustvmqueryordergetAPIRequest struct {
	model.Params
	// 阿里订单标编号
	_alitripOrderId string
}

// NewTaobaobustvmqueryordergetRequest 初始化TaobaobustvmqueryordergetAPIRequest对象
func NewTaobaobustvmqueryordergetRequest() *TaobaobustvmqueryordergetAPIRequest {
	return &TaobaobustvmqueryordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmqueryordergetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmqueryorder.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmqueryordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmqueryordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 阿里订单标编号
func (r *TaobaobustvmqueryordergetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmqueryordergetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}
