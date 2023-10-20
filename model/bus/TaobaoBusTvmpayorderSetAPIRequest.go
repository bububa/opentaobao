package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmpayorderSetAPIRequest 自助机条形码被动支付 API请求
// taobao.bus.tvmpayorder.set
//
// 汽车票线下自助机条形码支付
type TaobaoBusTvmpayorderSetAPIRequest struct {
	model.Params
	// 条形码认证码
	_alipayAuthCode string
	// 飞猪订单号
	_alitripOrderId string
}

// NewTaobaoBusTvmpayorderSetRequest 初始化TaobaoBusTvmpayorderSetAPIRequest对象
func NewTaobaoBusTvmpayorderSetRequest() *TaobaoBusTvmpayorderSetAPIRequest {
	return &TaobaoBusTvmpayorderSetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusTvmpayorderSetAPIRequest) Reset() {
	r._alipayAuthCode = ""
	r._alitripOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusTvmpayorderSetAPIRequest) GetApiMethodName() string {
	return "taobao.bus.tvmpayorder.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusTvmpayorderSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusTvmpayorderSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayAuthCode is AlipayAuthCode Setter
// 条形码认证码
func (r *TaobaoBusTvmpayorderSetAPIRequest) SetAlipayAuthCode(_alipayAuthCode string) error {
	r._alipayAuthCode = _alipayAuthCode
	r.Set("alipay_auth_code", _alipayAuthCode)
	return nil
}

// GetAlipayAuthCode AlipayAuthCode Getter
func (r TaobaoBusTvmpayorderSetAPIRequest) GetAlipayAuthCode() string {
	return r._alipayAuthCode
}

// SetAlitripOrderId is AlitripOrderId Setter
// 飞猪订单号
func (r *TaobaoBusTvmpayorderSetAPIRequest) SetAlitripOrderId(_alitripOrderId string) error {
	r._alitripOrderId = _alitripOrderId
	r.Set("alitrip_order_id", _alitripOrderId)
	return nil
}

// GetAlitripOrderId AlitripOrderId Getter
func (r TaobaoBusTvmpayorderSetAPIRequest) GetAlitripOrderId() string {
	return r._alitripOrderId
}

var poolTaobaoBusTvmpayorderSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusTvmpayorderSetRequest()
	},
}

// GetTaobaoBusTvmpayorderSetRequest 从 sync.Pool 获取 TaobaoBusTvmpayorderSetAPIRequest
func GetTaobaoBusTvmpayorderSetAPIRequest() *TaobaoBusTvmpayorderSetAPIRequest {
	return poolTaobaoBusTvmpayorderSetAPIRequest.Get().(*TaobaoBusTvmpayorderSetAPIRequest)
}

// ReleaseTaobaoBusTvmpayorderSetAPIRequest 将 TaobaoBusTvmpayorderSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusTvmpayorderSetAPIRequest(v *TaobaoBusTvmpayorderSetAPIRequest) {
	v.Reset()
	poolTaobaoBusTvmpayorderSetAPIRequest.Put(v)
}
