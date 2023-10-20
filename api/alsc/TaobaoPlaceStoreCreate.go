package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// TaobaoPlaceStoreCreate 商户门店创建接口
// taobao.place.store.create
//
// 用于商家创建线下门店
func TaobaoPlaceStoreCreate(clt *core.SDKClient, req *alsc.TaobaoPlaceStoreCreateAPIRequest, session string) (*alsc.TaobaoPlaceStoreCreateAPIResponse, error) {
	var resp alsc.TaobaoPlaceStoreCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
