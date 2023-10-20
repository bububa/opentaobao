package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobustvmcreateqrcodesetAPIRequest 自助机生成支付宝支付二维码 API请求
// taobao.bus.tvmcreateqrcode.set
//
// 用于汽车票线下自助机调用获取支付宝的二维码
type TaobaobustvmcreateqrcodesetAPIRequest struct {
	model.Params
	// 飞猪订单号
	_alitripOrderId string
	// 超时时间（分钟）
	_timeoutExpress int64
}

// NewTaobaobustvmcreateqrcodesetRequest 初始化TaobaobustvmcreateqrcodesetAPIRequest对象
func NewTaobaobustvmcreateqrcodesetRequest() *TaobaobustvmcreateqrcodesetAPIRequest {
	return &TaobaobustvmcreateqrcodesetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobustvmcreateqrcodesetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmcreateqrcode.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobustvmcreateqrcodesetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobustvmcreateqrcodesetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaobustvmcreateqrcodesetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaobustvmcreateqrcodesetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

// SetTimeoutExpress is TimeoutExpress Setter
// 超时时间（分钟）
func (r *TaobaobustvmcreateqrcodesetAPIRequest) SetTimeoutExpress(_timeoutExpress int64) error {
	r._timeoutExpress = _timeoutExpress
	r.Set("timeout_express", _timeoutExpress)
	return nil
}

// GetTimeoutExpress TimeoutExpress Getter
func (r TaobaobustvmcreateqrcodesetAPIRequest) GetTimeoutExpress() int64 {
	return r._timeoutExpress
}
