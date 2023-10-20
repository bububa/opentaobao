package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserGetAPIRequest 获取用户已开通消息 API请求
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
type TaobaoTmcUserGetAPIRequest struct {
	model.Params
	// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
	_fields string
	// 用户昵称
	_nick string
	// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaoTmcUserGetRequest 初始化TaobaoTmcUserGetAPIRequest对象
func NewTaobaoTmcUserGetRequest() *TaobaoTmcUserGetAPIRequest {
	return &TaobaoTmcUserGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcUserGetAPIRequest) Reset() {
	r._fields = ""
	r._nick = ""
	r._userPlatform = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserGetAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcUserGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。
func (r *TaobaoTmcUserGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoTmcUserGetAPIRequest) GetFields() string {
	return r._fields
}

// SetNick is Nick Setter
// 用户昵称
func (r *TaobaoTmcUserGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoTmcUserGetAPIRequest) GetNick() string {
	return r._nick
}

// SetUserPlatform is UserPlatform Setter
// 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcUserGetAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaoTmcUserGetAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}

var poolTaobaoTmcUserGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcUserGetRequest()
	},
}

// GetTaobaoTmcUserGetRequest 从 sync.Pool 获取 TaobaoTmcUserGetAPIRequest
func GetTaobaoTmcUserGetAPIRequest() *TaobaoTmcUserGetAPIRequest {
	return poolTaobaoTmcUserGetAPIRequest.Get().(*TaobaoTmcUserGetAPIRequest)
}

// ReleaseTaobaoTmcUserGetAPIRequest 将 TaobaoTmcUserGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcUserGetAPIRequest(v *TaobaoTmcUserGetAPIRequest) {
	v.Reset()
	poolTaobaoTmcUserGetAPIRequest.Put(v)
}
