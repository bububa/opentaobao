package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweitaofeedisrelationAPIRequest 是否关注 API请求
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoweitaofeedisrelationAPIRequest struct {
	model.Params
	// 要查询的公共账号的淘宝昵称
	_sellerNick string
	// 要查询的粉丝的淘宝昵称
	_fansNick string
}

// NewTaobaoweitaofeedisrelationRequest 初始化TaobaoweitaofeedisrelationAPIRequest对象
func NewTaobaoweitaofeedisrelationRequest() *TaobaoweitaofeedisrelationAPIRequest {
	return &TaobaoweitaofeedisrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweitaofeedisrelationAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.isrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweitaofeedisrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweitaofeedisrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 要查询的公共账号的淘宝昵称
func (r *TaobaoweitaofeedisrelationAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoweitaofeedisrelationAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetFansNick is FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoweitaofeedisrelationAPIRequest) SetFansNick(_fansNick string) error {
	r._fansNick = _fansNick
	r.Set("fans_nick", _fansNick)
	return nil
}

// GetFansNick FansNick Getter
func (r TaobaoweitaofeedisrelationAPIRequest) GetFansNick() string {
	return r._fansNick
}
