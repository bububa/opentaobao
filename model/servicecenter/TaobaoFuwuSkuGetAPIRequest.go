package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuSkuGetAPIRequest 获取内购服务及SKU详情 API请求
// taobao.fuwu.sku.get
//
// 通过服务code和用户nick，获取该服务对应的收费项目的sku信息，包括价格、可购买周期、用户能否购买等信息
type TaobaoFuwuSkuGetAPIRequest struct {
	model.Params
	// 服务code
	_articleCode string
	// 用户的淘宝nick
	_nick string
}

// NewTaobaoFuwuSkuGetRequest 初始化TaobaoFuwuSkuGetAPIRequest对象
func NewTaobaoFuwuSkuGetRequest() *TaobaoFuwuSkuGetAPIRequest {
	return &TaobaoFuwuSkuGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFuwuSkuGetAPIRequest) Reset() {
	r._articleCode = ""
	r._nick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSkuGetAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sku.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSkuGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFuwuSkuGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetArticleCode is ArticleCode Setter
// 服务code
func (r *TaobaoFuwuSkuGetAPIRequest) SetArticleCode(_articleCode string) error {
	r._articleCode = _articleCode
	r.Set("article_code", _articleCode)
	return nil
}

// GetArticleCode ArticleCode Getter
func (r TaobaoFuwuSkuGetAPIRequest) GetArticleCode() string {
	return r._articleCode
}

// SetNick is Nick Setter
// 用户的淘宝nick
func (r *TaobaoFuwuSkuGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoFuwuSkuGetAPIRequest) GetNick() string {
	return r._nick
}

var poolTaobaoFuwuSkuGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFuwuSkuGetRequest()
	},
}

// GetTaobaoFuwuSkuGetRequest 从 sync.Pool 获取 TaobaoFuwuSkuGetAPIRequest
func GetTaobaoFuwuSkuGetAPIRequest() *TaobaoFuwuSkuGetAPIRequest {
	return poolTaobaoFuwuSkuGetAPIRequest.Get().(*TaobaoFuwuSkuGetAPIRequest)
}

// ReleaseTaobaoFuwuSkuGetAPIRequest 将 TaobaoFuwuSkuGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFuwuSkuGetAPIRequest(v *TaobaoFuwuSkuGetAPIRequest) {
	v.Reset()
	poolTaobaoFuwuSkuGetAPIRequest.Put(v)
}
