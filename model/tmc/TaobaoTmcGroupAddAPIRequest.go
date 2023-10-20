package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcgroupaddAPIRequest 为已开通用户添加用户分组 API请求
// taobao.tmc.group.add
//
// 为已开通用户添加用户分组，授权消息使用
type TaobaotmcgroupaddAPIRequest struct {
	model.Params
	// 用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户
	_nicks []string
	// 分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。
	_groupName string
	// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaotmcgroupaddRequest 初始化TaobaotmcgroupaddAPIRequest对象
func NewTaobaotmcgroupaddRequest() *TaobaotmcgroupaddAPIRequest {
	return &TaobaotmcgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotmcgroupaddAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.group.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotmcgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotmcgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNicks is Nicks Setter
// 用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户
func (r *TaobaotmcgroupaddAPIRequest) SetNicks(_nicks []string) error {
	r._nicks = _nicks
	r.Set("nicks", _nicks)
	return nil
}

// GetNicks Nicks Getter
func (r TaobaotmcgroupaddAPIRequest) GetNicks() []string {
	return r._nicks
}

// SetGroupName is GroupName Setter
// 分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。
func (r *TaobaotmcgroupaddAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaotmcgroupaddAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetUserPlatform is UserPlatform Setter
// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaotmcgroupaddAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaotmcgroupaddAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}
