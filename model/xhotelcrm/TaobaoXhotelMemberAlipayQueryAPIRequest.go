package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelmemberalipayqueryAPIRequest 希尔顿会员查询 API请求
// taobao.xhotel.member.alipay.query
//
// 希尔顿会员查询
type TaobaoxhotelmemberalipayqueryAPIRequest struct {
	model.Params
	// 渠道
	_fpt string
	// 支付宝id
	_alipayUserId int64
}

// NewTaobaoxhotelmemberalipayqueryRequest 初始化TaobaoxhotelmemberalipayqueryAPIRequest对象
func NewTaobaoxhotelmemberalipayqueryRequest() *TaobaoxhotelmemberalipayqueryAPIRequest {
	return &TaobaoxhotelmemberalipayqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelmemberalipayqueryAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.member.alipay.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelmemberalipayqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelmemberalipayqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFpt is Fpt Setter
// 渠道
func (r *TaobaoxhotelmemberalipayqueryAPIRequest) SetFpt(_fpt string) error {
	r._fpt = _fpt
	r.Set("fpt", _fpt)
	return nil
}

// GetFpt Fpt Getter
func (r TaobaoxhotelmemberalipayqueryAPIRequest) GetFpt() string {
	return r._fpt
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝id
func (r *TaobaoxhotelmemberalipayqueryAPIRequest) SetAlipayUserId(_alipayUserId int64) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r TaobaoxhotelmemberalipayqueryAPIRequest) GetAlipayUserId() int64 {
	return r._alipayUserId
}
