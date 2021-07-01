package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJmsGroupDeleteAPIRequest
删除ONS分组 API请求
taobao.jushita.jms.group.delete

删除ONS分组 */
type TaobaoJushitaJmsGroupDeleteAPIRequest struct {
	model.Params
	// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
	_groupName string
	// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
	_nicks []string
	// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户
	_userPlatform string
}

// NewTaobaoJushitaJmsGroupDeleteRequest 初始化TaobaoJushitaJmsGroupDeleteAPIRequest对象
func NewTaobaoJushitaJmsGroupDeleteRequest() *TaobaoJushitaJmsGroupDeleteAPIRequest {
	return &TaobaoJushitaJmsGroupDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsGroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsGroupDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupName Setter
// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
func (r *TaobaoJushitaJmsGroupDeleteAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// Get GroupName Getter
func (r TaobaoJushitaJmsGroupDeleteAPIRequest) GetGroupName() string {
	return r._groupName
}

// Set is Nicks Setter
// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
func (r *TaobaoJushitaJmsGroupDeleteAPIRequest) SetNicks(_nicks []string) error {
	r._nicks = _nicks
	r.Set("nicks", _nicks)
	return nil
}

// Get Nicks Getter
func (r TaobaoJushitaJmsGroupDeleteAPIRequest) GetNicks() []string {
	return r._nicks
}

// Set is UserPlatform Setter
// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户
func (r *TaobaoJushitaJmsGroupDeleteAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// Get UserPlatform Getter
func (r TaobaoJushitaJmsGroupDeleteAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}
