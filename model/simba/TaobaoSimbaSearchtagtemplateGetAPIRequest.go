package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasearchtagtemplategetAPIRequest 获取搜索人群TOP用户可添加人群信息 API请求
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
type TaobaosimbasearchtagtemplategetAPIRequest struct {
	model.Params
	// 子帐号nick
	_subNick string
	// 被操作者的淘宝昵称
	_nick string
}

// NewTaobaosimbasearchtagtemplategetRequest 初始化TaobaosimbasearchtagtemplategetAPIRequest对象
func NewTaobaosimbasearchtagtemplategetRequest() *TaobaosimbasearchtagtemplategetAPIRequest {
	return &TaobaosimbasearchtagtemplategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasearchtagtemplategetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.searchtagtemplate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasearchtagtemplategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasearchtagtemplategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSubNick is SubNick Setter
// 子帐号nick
func (r *TaobaosimbasearchtagtemplategetAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaosimbasearchtagtemplategetAPIRequest) GetSubNick() string {
	return r._subNick
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbasearchtagtemplategetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbasearchtagtemplategetAPIRequest) GetNick() string {
	return r._nick
}
