package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStoreDelete 线下门店删除
// taobao.place.store.delete
//
// 用于商家删除线下门店
func TaobaoPlaceStoreDelete(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreDeleteAPIRequest, resp *alsc.TaobaoPlaceStoreDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
