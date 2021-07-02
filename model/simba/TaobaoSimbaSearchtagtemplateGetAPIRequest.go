package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSearchtagtemplateGetAPIRequest 获取搜索人群TOP用户可添加人群信息 API请求
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
type TaobaoSimbaSearchtagtemplateGetAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 子帐号nick
	_subNick string
}

// NewTaobaoSimbaSearchtagtemplateGetRequest 初始化TaobaoSimbaSearchtagtemplateGetAPIRequest对象
func NewTaobaoSimbaSearchtagtemplateGetRequest() *TaobaoSimbaSearchtagtemplateGetAPIRequest {
	return &TaobaoSimbaSearchtagtemplateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.searchtagtemplate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSearchtagtemplateGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetNick() string {
	return r._nick
}

// SetSubNick is SubNick Setter
// 子帐号nick
func (r *TaobaoSimbaSearchtagtemplateGetAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetSubNick() string {
	return r._subNick
}
