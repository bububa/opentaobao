package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalihealthdrugusershopgetAPIRequest 根据用户id获取店铺id API请求
// taobao.alihealth.drug.user.shop.get
//
// 提供给千牛智能客服，获取用户当前咨询的店铺ID
type TaobaoalihealthdrugusershopgetAPIRequest struct {
	model.Params
	// 用户昵称
	_userNick string
}

// NewTaobaoalihealthdrugusershopgetRequest 初始化TaobaoalihealthdrugusershopgetAPIRequest对象
func NewTaobaoalihealthdrugusershopgetRequest() *TaobaoalihealthdrugusershopgetAPIRequest {
	return &TaobaoalihealthdrugusershopgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalihealthdrugusershopgetAPIRequest) GetApiMethodName() string {
	return "taobao.alihealth.drug.user.shop.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalihealthdrugusershopgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalihealthdrugusershopgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 用户昵称
func (r *TaobaoalihealthdrugusershopgetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoalihealthdrugusershopgetAPIRequest) GetUserNick() string {
	return r._userNick
}
