package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasSubscribeGetAPIRequest
订购关系查询 API请求
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVasSubscribeGetAPIRequest) GetApiMethodName() string {
	return "taobao.vas.subscribe.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVasSubscribeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ArticleCode Setter
// 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码
func (r *TaobaoVasSubscribeGetAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// Get ArticleCode Getter
func (r TaobaoVasSubscribeGetAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// Set is Nick Setter
// 淘宝会员名
func (r *TaobaoVasSubscribeGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoVasSubscribeGetAPIRequest) GetNick() string {
	return r._nick
}
