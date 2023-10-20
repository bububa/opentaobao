package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoStoreFollowurlGet 获取店铺关注URL
// taobao.store.followurl.get
//
// 获取关注店铺的URL
func TaobaoStoreFollowurlGet(clt *core.SDKClient, req *shop.TaobaoStoreFollowurlGetAPIRequest, resp *shop.TaobaoStoreFollowurlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
