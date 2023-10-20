package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Taobaostorefollowurlget 获取店铺关注URL
// taobao.store.followurl.get
//
// 获取关注店铺的URL
func Taobaostorefollowurlget(clt *core.SDKClient, req *shop.TaobaostorefollowurlgetAPIRequest, session string) (*shop.TaobaostorefollowurlgetAPIResponse, error) {
	var resp shop.TaobaostorefollowurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
