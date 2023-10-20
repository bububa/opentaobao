package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaostorefollowurlgetAPIRequest 获取店铺关注URL API请求
// taobao.store.followurl.get
//
// 获取关注店铺的URL
type TaobaostorefollowurlgetAPIRequest struct {
	model.Params
	// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
	_callbackUrl string
	// 商家nick
	_userNick string
	// 商家ID
	_userId int64
}

// NewTaobaostorefollowurlgetRequest 初始化TaobaostorefollowurlgetAPIRequest对象
func NewTaobaostorefollowurlgetRequest() *TaobaostorefollowurlgetAPIRequest {
	return &TaobaostorefollowurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaostorefollowurlgetAPIRequest) GetApiMethodName() string {
	return "taobao.store.followurl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaostorefollowurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaostorefollowurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCallbackUrl is CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
func (r *TaobaostorefollowurlgetAPIRequest) SetCallbackUrl(_callbackUrl string) error {
	r._callbackUrl = _callbackUrl
	r.Set("callback_url", _callbackUrl)
	return nil
}

// GetCallbackUrl CallbackUrl Getter
func (r TaobaostorefollowurlgetAPIRequest) GetCallbackUrl() string {
	return r._callbackUrl
}

// SetUserNick is UserNick Setter
// 商家nick
func (r *TaobaostorefollowurlgetAPIRequest) SetUserNick(_userNick string) error {
	r._userNick = _userNick
	r.Set("user_nick", _userNick)
	return nil
}

// GetUserNick UserNick Getter
func (r TaobaostorefollowurlgetAPIRequest) GetUserNick() string {
	return r._userNick
}

// SetUserId is UserId Setter
// 商家ID
func (r *TaobaostorefollowurlgetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaostorefollowurlgetAPIRequest) GetUserId() int64 {
	return r._userId
}
