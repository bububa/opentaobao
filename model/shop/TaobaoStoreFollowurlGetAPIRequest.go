package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoStoreFollowurlGetAPIRequest
获取店铺关注URL API请求
taobao.store.followurl.get

获取关注店铺的URL */
type TaobaoStoreFollowurlGetAPIRequest struct {
	model.Params
	// 关注完成后的回调地址,需要是EWS地址。如果不设置，会跳转到店铺首页
	_callbackUrl string
	// 商家nick
	_userNick string
	// 商家ID
	_userId int64
}

// New
