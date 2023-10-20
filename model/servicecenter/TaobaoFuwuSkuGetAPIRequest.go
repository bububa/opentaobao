package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofuwuskugetAPIRequest 获取内购服务及SKU详情 API请求
// taobao.fuwu.sku.get
//
// 通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
type TaobaofuwuskugetAPIRequest struct {
	model.Params
	// 服务code
	_articleCode string
	// 用户的淘宝nick
	_nick string
}

// NewTaobaofuwuskugetRequest 初始化TaobaofuwuskugetAPIRequest对象
func NewTaobaofuwuskugetRequest() *TaobaofuwuskugetAPIRequest {
	return &TaobaofuwuskugetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofuwuskugetAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofuwuskugetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofuwuskugetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArticleCode is ArticleCode Setter
// 服务code
func (r *TaobaofuwuskugetAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaofuwuskugetAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetNick is Nick Setter
// 用户的淘宝nick
func (r *TaobaofuwuskugetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaofuwuskugetAPIRequest) GetNick() string {
	return r._nick
}
