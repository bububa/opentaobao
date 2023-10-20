package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoappstoresubscribegetAPIRequest 查询appstore应用订购关系 API请求
// taobao.appstore.subscribe.get
//
// 查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
type TaobaoappstoresubscribegetAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
}

// NewTaobaoappstoresubscribegetRequest 初始化TaobaoappstoresubscribegetAPIRequest对象
func NewTaobaoappstoresubscribegetRequest() *TaobaoappstoresubscribegetAPIRequest {
	return &TaobaoappstoresubscribegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoappstoresubscribegetAPIRequest) GetApiMethodName() string {
	return "taobao.appstore.subscribe.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoappstoresubscribegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoappstoresubscribegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoappstoresubscribegetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoappstoresubscribegetAPIRequest) GetNick() string {
	return r._nick
}
