package shop

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoStoreFollowurlGetAPIRequest) Reset() {
	r._callbackUrl = ""
	r._userNick = ""
	r._userId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoStoreFollowurlGetAPIRequest) GetApiMethodName() string {
	return "taobao.store.followurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoStoreFollowurlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoStoreFollowurlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackUrl is CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
func (r *TaobaoStoreFollowurlGetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetUserNick is UserNick Setter
// 商家nick
func (r *TaobaoStoreFollowurlGetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetUserId is UserId Setter
// 商家ID
func (r *TaobaoStoreFollowurlGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoStoreFollowurlGetAPIRequest) GetUserId() int64 {
	return r._userId
}

var poolTaobaoStoreFollowurlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoStoreFollowurlGetRequest()
	},
}

// GetTaobaoStoreFollowurlGetRequest 从 sync.Pool 获取 TaobaoStoreFollowurlGetAPIRequest
func GetTaobaoStoreFollowurlGetAPIRequest() *TaobaoStoreFollowurlGetAPIRequest {
	return poolTaobaoStoreFollowurlGetAPIRequest.Get().(*TaobaoStoreFollowurlGetAPIRequest)
}

// ReleaseTaobaoStoreFollowurlGetAPIRequest 将 TaobaoStoreFollowurlGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoStoreFollowurlGetAPIRequest(v *TaobaoStoreFollowurlGetAPIRequest) {
	v.Reset()
	poolTaobaoStoreFollowurlGetAPIRequest.Put(v)
}
