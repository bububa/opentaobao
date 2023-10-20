package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowalipayisbindingtbaccountAPIRequest 判断支付宝用户是否绑定淘宝账号 API请求
// alibaba.aliqin.flow.alipay.isbindingtbaccount
//
// 判断支付宝用户是否绑定淘宝账号
type AlibabaaliqinflowalipayisbindingtbaccountAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayId string
}

// NewAlibabaaliqinflowalipayisbindingtbaccountRequest 初始化AlibabaaliqinflowalipayisbindingtbaccountAPIRequest对象
func NewAlibabaaliqinflowalipayisbindingtbaccountRequest() *AlibabaaliqinflowalipayisbindingtbaccountAPIRequest {
	return &AlibabaaliqinflowalipayisbindingtbaccountAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowalipayisbindingtbaccountAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.alipay.isbindingtbaccount"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowalipayisbindingtbaccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowalipayisbindingtbaccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlipayId is AlipayId Setter
// 支付宝ID
func (r *AlibabaaliqinflowalipayisbindingtbaccountAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabaaliqinflowalipayisbindingtbaccountAPIRequest) GetAlipayId() string {
	return r._alipayId
}
