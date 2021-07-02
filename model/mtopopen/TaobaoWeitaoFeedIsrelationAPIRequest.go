package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedIsrelationAPIRequest 是否关注 API请求
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoWeitaoFeedIsrelationAPIRequest struct {
	model.Params
	// 要查询的粉丝的淘宝昵称
	_fansNick string
	// 要查询的公共账号的淘宝昵称
	_sellerNick string
}

// NewTaobaoWeitaoFeedIsrelationRequest 初始化TaobaoWeitaoFeedIsrelationAPIRequest对象
func NewTaobaoWeitaoFeedIsrelationRequest() *TaobaoWeitaoFeedIsrelationAPIRequest {
	return &TaobaoWeitaoFeedIsrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.isrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationAPIRequest) SetFansNick(_fansNick string) error {
	r._fansNick = _fansNick
	r.Set("fans_nick", _fansNick)
	return nil
}

// Get FansNick Getter
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetFansNick() string {
	return r._fansNick
}

// Set is SellerNick Setter
// 要查询的公共账号的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// Get SellerNick Getter
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetSellerNick() string {
	return r._sellerNick
}
