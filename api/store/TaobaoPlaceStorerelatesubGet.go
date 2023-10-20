package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStorerelatesubGet 门店和子门店关系查找
// taobao.place.storerelatesub.get
//
// 门店和子门店关系查找
func TaobaoPlaceStorerelatesubGet(clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubGetAPIRequest, resp *store.TaobaoPlaceStorerelatesubGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
