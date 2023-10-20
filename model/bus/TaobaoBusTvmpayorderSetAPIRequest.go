package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmpayordersetAPIRequest 自助机条形码被动支付 API请求
// taobao.bus.tvmpayorder.set
//
// 汽车票线下自助机条形码支付
type TaobaobustvmpayordersetAPIRequest struct {
	model.Params
	// 条形码认证码
	_alipayAuthCode string
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaobustvmpayordersetRequest 初始化TaobaobustvmpayordersetAPIRequest对象
func NewTaobaobustvmpayordersetRequest() *TaobaobustvmpayordersetAPIRequest {
	return &TaobaobustvmpayordersetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmpayordersetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmpayorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmpayordersetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmpayordersetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayAuthCode is AlipayAuthCode Setter
// 条形码认证码
func (r *TaobaobustvmpayordersetAPIRequest) SetAlipayAuthCode(_alipayAuthCode string) error {
	r._alipayAuthCode = _alipayAuthCode
	r.Set("alipay_auth_code", _alipayAuthCode)
	return nil
}

// GetAlipayAuthCode AlipayAuthCode Getter
func (r TaobaobustvmpayordersetAPIRequest) GetAlipayAuthCode() string {
	return r._alipayAuthCode
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobustvmpayordersetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmpayordersetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}
