package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoregroupUpdate 门店库修改基本信息
// taobao.place.storegroup.update
//
// 门店库修改基本信息
func TaobaoPlaceStoregroupUpdate(clt *core.SDKClient, req *store.TaobaoPlaceStoregroupUpdateAPIRequest, resp *store.TaobaoPlaceStoregroupUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
