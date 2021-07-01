package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* TaobaoPlaceStoreDelete
线下门店删除
taobao.place.store.delete

用于商家删除线下门店 */
func TaobaoPlaceStoreDelete(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreDeleteAPIRequest, session string) (*alsc.TaobaoPlaceStoreDeleteAPIResponse, error) {
	var resp alsc.TaobaoPlaceStoreDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
