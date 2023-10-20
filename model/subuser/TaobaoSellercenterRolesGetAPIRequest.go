package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosellercenterrolesgetAPIRequest 获取指定卖家的角色列表 API请求
// taobao.sellercenter.roles.get
//
// 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。
type TaobaosellercenterrolesgetAPIRequest struct {
	model.Params
	// 卖家昵称(只允许查询自己的信息：当前登陆者)
	_nick string
}

// NewTaobaosellercenterrolesgetRequest 初始化TaobaosellercenterrolesgetAPIRequest对象
func NewTaobaosellercenterrolesgetRequest() *TaobaosellercenterrolesgetAPIRequest {
	return &TaobaosellercenterrolesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosellercenterrolesgetAPIRequest) GetApiMethodName() string {
	return "taobao.sellercenter.roles.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosellercenterrolesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosellercenterrolesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 卖家昵称(只允许查询自己的信息：当前登陆者)
func (r *TaobaosellercenterrolesgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosellercenterrolesgetAPIRequest) GetNick() string {
	return r._nick
}
