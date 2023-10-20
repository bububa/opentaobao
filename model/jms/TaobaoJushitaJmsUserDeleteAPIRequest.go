package jms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajmsuserdeleteAPIRequest 删除ONS消息同步用户 API请求
// taobao.jushita.jms.user.delete
//
// 删除ONS消息同步用户，删除后用户的消息将不会推送到聚石塔的ONS中
type TaobaojushitajmsuserdeleteAPIRequest struct {
	model.Params
	// 需要停止同步消息的用户nick
	_userNick string
}

// NewTaobaojushitajmsuserdeleteRequest 初始化TaobaojushitajmsuserdeleteAPIRequest对象
func NewTaobaojushitajmsuserdeleteRequest() *TaobaojushitajmsuserdeleteAPIRequest {
	return &TaobaojushitajmsuserdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajmsuserdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jms.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajmsuserdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajmsuserdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserNick is UserNick Setter
// 需要停止同步消息的用户nick
func (r *TaobaojushitajmsuserdeleteAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaojushitajmsuserdeleteAPIRequest) GetUserNick() string {
	return r._userNick
}
