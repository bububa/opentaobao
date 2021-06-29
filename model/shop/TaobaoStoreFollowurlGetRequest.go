package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取店铺关注URL API请求
taobao.store.followurl.get

获取关注店铺的URL
*/
type TaobaoStoreFollowurlGetRequest struct {
    model.Params
    // 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
    _callbackUrl   string
    // 商家nick
    _userNick   string
    // 商家ID
    _userId   int64
}

// 初始化TaobaoStoreFollowurlGetRequest对象
func NewTaobaoStoreFollowurlGetRequest() *TaobaoStoreFollowurlGetRequest{
    return &TaobaoStoreFollowurlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoStoreFollowurlGetRequest) GetApiMethodName() string {
    return "taobao.store.followurl.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoStoreFollowurlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CallbackUrl Setter
// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
func (r *TaobaoStoreFollowurlGetRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r TaobaoStoreFollowurlGetRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// UserNick Setter
// 商家nick
func (r *TaobaoStoreFollowurlGetRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r TaobaoStoreFollowurlGetRequest) GetUserNick() string {
    return r._userNick
}
// UserId Setter
// 商家ID
func (r *TaobaoStoreFollowurlGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoStoreFollowurlGetRequest) GetUserId() int64 {
    return r._userId
}
