package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoStoreFollowurlGetAPIRequest 获取店铺关注URL API请求
// taobao.store.followurl.get
//
// 获取关注店铺的URL
type TaobaoStoreFollowurlGetAPIRequest struct {
	model.Params
	// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
	_callbackUrl string
	// 商家nick
	_userNick string
	// 商家ID
	_userId int64
}

// NewTaobaoStoreFollowurlGetRequest 初始化TaobaoStoreFollowurlGetAPIRequest对象
func NewTaobaoStoreFollowurlGetRequest() *TaobaoStoreFollowurlGetAPIRequest {
	return &TaobaoStoreFollowurlGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoStoreFollowurlGetAPIRequest) GetApiMethodName() string {
	return "taobao.store.followurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoStoreFollowurlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
func (r *TaobaoStoreFollowurlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// Get CallbackUrl Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// Set is UserNick Setter
// 商家nick
func (r *TaobaoStoreFollowurlGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// Get UserNick Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetUserNick() string {
	return r._userNick
}

// Set is UserId Setter
// 商家ID
func (r *TaobaoStoreFollowurlGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetUserId() int64 {
	return r._userId
}
