package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVasSubscribeGetAPIRequest 订购关系查询 API请求
// taobao.vas.subscribe.get
//
// 用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
type TaobaoVasSubscribeGetAPIRequest struct {
	model.Params
	// 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
	_articleCode string
	// 淘宝会员名
	_nick string
}

// NewTaobaoVasSubscribeGetRequest 初始化TaobaoVasSubscribeGetAPIRequest对象
func NewTaobaoVasSubscribeGetRequest() *TaobaoVasSubscribeGetAPIRequest {
	return &TaobaoVasSubscribeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVasSubscribeGetAPIRequest) Reset() {
	r._articleCode = ""
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasSubscribeGetAPIRequest) GetApiMethodName() string {
	return "taobao.vas.subscribe.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasSubscribeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVasSubscribeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArticleCode is ArticleCode Setter
// 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
func (r *TaobaoVasSubscribeGetAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaoVasSubscribeGetAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetNick is Nick Setter
// 淘宝会员名
func (r *TaobaoVasSubscribeGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoVasSubscribeGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoVasSubscribeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVasSubscribeGetRequest()
	},
}

// GetTaobaoVasSubscribeGetRequest 从 sync.Pool 获取 TaobaoVasSubscribeGetAPIRequest
func GetTaobaoVasSubscribeGetAPIRequest() *TaobaoVasSubscribeGetAPIRequest {
	return poolTaobaoVasSubscribeGetAPIRequest.Get().(*TaobaoVasSubscribeGetAPIRequest)
}

// ReleaseTaobaoVasSubscribeGetAPIRequest 将 TaobaoVasSubscribeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoVasSubscribeGetAPIRequest(v *TaobaoVasSubscribeGetAPIRequest) {
	v.Reset()
	poolTaobaoVasSubscribeGetAPIRequest.Put(v)
}
