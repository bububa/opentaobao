package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcGroupDeleteAPIRequest 删除指定的分组或分组下的用户 API请求
// taobao.tmc.group.delete
//
// 删除指定的分组或分组下的用户，授权消息使用
type TaobaoTmcGroupDeleteAPIRequest struct {
	model.Params
	// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
	_nicks []string
	// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
	_groupName string
	// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
	_userPlatform string
}

// NewTaobaoTmcGroupDeleteRequest 初始化TaobaoTmcGroupDeleteAPIRequest对象
func NewTaobaoTmcGroupDeleteRequest() *TaobaoTmcGroupDeleteAPIRequest {
	return &TaobaoTmcGroupDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcGroupDeleteAPIRequest) Reset() {
	r._nicks = r._nicks[:0]
	r._groupName = ""
	r._userPlatform = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcGroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.group.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcGroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcGroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNicks is Nicks Setter
// 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组
func (r *TaobaoTmcGroupDeleteAPIRequest) SetNicks(_nicks []string) error {
	r._nicks = _nicks
	r.Set("nicks", _nicks)
	return nil
}

// GetNicks Nicks Getter
func (r TaobaoTmcGroupDeleteAPIRequest) GetNicks() []string {
	return r._nicks
}

// SetGroupName is GroupName Setter
// 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。
func (r *TaobaoTmcGroupDeleteAPIRequest) SetGroupName(_groupName string) error {
	r._groupName = _groupName
	r.Set("group_name", _groupName)
	return nil
}

// GetGroupName GroupName Getter
func (r TaobaoTmcGroupDeleteAPIRequest) GetGroupName() string {
	return r._groupName
}

// SetUserPlatform is UserPlatform Setter
// 用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
func (r *TaobaoTmcGroupDeleteAPIRequest) SetUserPlatform(_userPlatform string) error {
	r._userPlatform = _userPlatform
	r.Set("user_platform", _userPlatform)
	return nil
}

// GetUserPlatform UserPlatform Getter
func (r TaobaoTmcGroupDeleteAPIRequest) GetUserPlatform() string {
	return r._userPlatform
}

var poolTaobaoTmcGroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcGroupDeleteRequest()
	},
}

// GetTaobaoTmcGroupDeleteRequest 从 sync.Pool 获取 TaobaoTmcGroupDeleteAPIRequest
func GetTaobaoTmcGroupDeleteAPIRequest() *TaobaoTmcGroupDeleteAPIRequest {
	return poolTaobaoTmcGroupDeleteAPIRequest.Get().(*TaobaoTmcGroupDeleteAPIRequest)
}

// ReleaseTaobaoTmcGroupDeleteAPIRequest 将 TaobaoTmcGroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcGroupDeleteAPIRequest(v *TaobaoTmcGroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoTmcGroupDeleteAPIRequest.Put(v)
}
