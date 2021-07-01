package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* TaobaoPlaceStoreModify
商家修改线下门店
taobao.place.store.modify

用于商家修改线下门店信息 */
func TaobaoPlaceStoreModify(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreModifyAPIRequest, session string) (*alsc.TaobaoPlaceStoreModifyAPIResponse, error) {
	var resp alsc.TaobaoPlaceStoreModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
