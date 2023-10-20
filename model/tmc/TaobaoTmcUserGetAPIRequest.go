package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcusergetAPIRequest 获取用户已开通消息 API请求
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
type TaobaotmcusergetAPIRequest struct {
	model.Params
	// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
	_fields string
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaotmcusergetRequest 初始化TaobaotmcusergetAPIRequest对象
func NewTaobaotmcusergetRequest() *TaobaotmcusergetAPIRequest {
	return &TaobaotmcusergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcusergetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcusergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcusergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
func (r *TaobaotmcusergetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaotmcusergetAPIRequest) GetFields() string {
	return r._fields
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaotmcusergetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaotmcusergetAPIRequest) GetNick() string {
	return r._nick
}

// SetUserPlatform is UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaotmcusergetAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaotmcusergetAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}
