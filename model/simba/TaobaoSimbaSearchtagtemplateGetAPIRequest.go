package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSearchtagtemplateGetAPIRequest 获取搜索人群TOP用户可添加人群信息 API请求
// taobao.simba.searchtagtemplate.get
//
// 获取搜索人群用户可添加人群信息
type TaobaoSimbaSearchtagtemplateGetAPIRequest struct {
	model.Params
	// 子帐号nick
	_subNick string
	// 被操作者的淘宝昵称
	_nick string
}

// NewTaobaoSimbaSearchtagtemplateGetRequest 初始化TaobaoSimbaSearchtagtemplateGetAPIRequest对象
func NewTaobaoSimbaSearchtagtemplateGetRequest() *TaobaoSimbaSearchtagtemplateGetAPIRequest {
	return &TaobaoSimbaSearchtagtemplateGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSearchtagtemplateGetAPIRequest) Reset() {
	r._subNick = ""
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.searchtagtemplate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSearchtagtemplateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoSimbaSearchtagtemplateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSearchtagtemplateGetRequest()
	},
}

// GetTaobaoSimbaSearchtagtemplateGetRequest 从 sync.Pool 获取 TaobaoSimbaSearchtagtemplateGetAPIRequest
func GetTaobaoSimbaSearchtagtemplateGetAPIRequest() *TaobaoSimbaSearchtagtemplateGetAPIRequest {
	return poolTaobaoSimbaSearchtagtemplateGetAPIRequest.Get().(*TaobaoSimbaSearchtagtemplateGetAPIRequest)
}

// ReleaseTaobaoSimbaSearchtagtemplateGetAPIRequest 将 TaobaoSimbaSearchtagtemplateGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSearchtagtemplateGetAPIRequest(v *TaobaoSimbaSearchtagtemplateGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSearchtagtemplateGetAPIRequest.Put(v)
}
