package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStorerelatesubDelete 门店和子门店关系删除
// taobao.place.storerelatesub.delete
//
// 门店和子门店关系删除
func TaobaoPlaceStorerelatesubDelete(clt *core.SDKClient, req *store.TaobaoPlaceStorerelatesubDeleteAPIRequest, resp *store.TaobaoPlaceStorerelatesubDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
