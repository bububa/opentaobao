package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofilmaccountphonequeryAPIRequest 根据手机查询匹配账号列表 API请求
// taobao.film.account.phone.query
//
// 根据手机号查询匹配的账号列表
type TaobaofilmaccountphonequeryAPIRequest struct {
	model.Params
	// 手机号
	_phone string
}

// NewTaobaofilmaccountphonequeryRequest 初始化TaobaofilmaccountphonequeryAPIRequest对象
func NewTaobaofilmaccountphonequeryRequest() *TaobaofilmaccountphonequeryAPIRequest {
	return &TaobaofilmaccountphonequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofilmaccountphonequeryAPIRequest) GetApiMethodName() string {
	return "taobao.film.account.phone.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofilmaccountphonequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofilmaccountphonequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhone is Phone Setter
// 手机号
func (r *TaobaofilmaccountphonequeryAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaofilmaccountphonequeryAPIRequest) GetPhone() string {
	return r._phone
}
