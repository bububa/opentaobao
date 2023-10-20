package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojushitajdpuserdeleteAPIRequest 删除数据推送用户 API请求
// taobao.jushita.jdp.user.delete
//
// 删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。
type TaobaojushitajdpuserdeleteAPIRequest struct {
	model.Params
	// 要删除用户的昵称
	_nick string
}

// NewTaobaojushitajdpuserdeleteRequest 初始化TaobaojushitajdpuserdeleteAPIRequest对象
func NewTaobaojushitajdpuserdeleteRequest() *TaobaojushitajdpuserdeleteAPIRequest {
	return &TaobaojushitajdpuserdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojushitajdpuserdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojushitajdpuserdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojushitajdpuserdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 要删除用户的昵称
func (r *TaobaojushitajdpuserdeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaojushitajdpuserdeleteAPIRequest) GetNick() string {
	return r._nick
}
