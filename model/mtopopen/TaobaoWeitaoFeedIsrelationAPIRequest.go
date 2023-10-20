package mtopopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeitaoFeedIsrelationAPIRequest 是否关注 API请求
// taobao.weitao.feed.isrelation
//
// 判断用户是否关注对应的公共账号
type TaobaoWeitaoFeedIsrelationAPIRequest struct {
	model.Params
	// 要查询的公共账号的淘宝昵称
	_sellerNick string
	// 要查询的粉丝的淘宝昵称
	_fansNick string
}

// NewTaobaoWeitaoFeedIsrelationRequest 初始化TaobaoWeitaoFeedIsrelationAPIRequest对象
func NewTaobaoWeitaoFeedIsrelationRequest() *TaobaoWeitaoFeedIsrelationAPIRequest {
	return &TaobaoWeitaoFeedIsrelationAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeitaoFeedIsrelationAPIRequest) Reset() {
	r._sellerNick = ""
	r._fansNick = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetApiMethodName() string {
	return "taobao.weitao.feed.isrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 要查询的公共账号的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetFansNick is FansNick Setter
// 要查询的粉丝的淘宝昵称
func (r *TaobaoWeitaoFeedIsrelationAPIRequest) SetFansNick(_fansNick string) error {
	r._fansNick = _fansNick
	r.Set("fans_nick", _fansNick)
	return nil
}

// GetFansNick FansNick Getter
func (r TaobaoWeitaoFeedIsrelationAPIRequest) GetFansNick() string {
	return r._fansNick
}

var poolTaobaoWeitaoFeedIsrelationAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeitaoFeedIsrelationRequest()
	},
}

// GetTaobaoWeitaoFeedIsrelationRequest 从 sync.Pool 获取 TaobaoWeitaoFeedIsrelationAPIRequest
func GetTaobaoWeitaoFeedIsrelationAPIRequest() *TaobaoWeitaoFeedIsrelationAPIRequest {
	return poolTaobaoWeitaoFeedIsrelationAPIRequest.Get().(*TaobaoWeitaoFeedIsrelationAPIRequest)
}

// ReleaseTaobaoWeitaoFeedIsrelationAPIRequest 将 TaobaoWeitaoFeedIsrelationAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeitaoFeedIsrelationAPIRequest(v *TaobaoWeitaoFeedIsrelationAPIRequest) {
	v.Reset()
	poolTaobaoWeitaoFeedIsrelationAPIRequest.Put(v)
}
