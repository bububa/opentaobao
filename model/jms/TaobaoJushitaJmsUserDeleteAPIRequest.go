package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsUserDeleteAPIRequest 删除ONS消息同步用户 API请求
// taobao.jushita.jms.user.delete
//
// 删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
type TaobaoJushitaJmsUserDeleteAPIRequest struct {
	model.Params
	// 需要停止同步消息的用户nick
	_userNick string
}

// NewTaobaoJushitaJmsUserDeleteRequest 初始化TaobaoJushitaJmsUserDeleteAPIRequest对象
func NewTaobaoJushitaJmsUserDeleteRequest() *TaobaoJushitaJmsUserDeleteAPIRequest {
	return &TaobaoJushitaJmsUserDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJushitaJmsUserDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJushitaJmsUserDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserNick Setter
// 需要停止同步消息的用户nick
func (r *TaobaoJushitaJmsUserDeleteAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r TaobaoJushitaJmsUserDeleteAPIRequest) GetUserNick() string {
	return r._userNick
}
