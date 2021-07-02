package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoregroupDelete 删除门店库
// taobao.place.storegroup.delete
//
// 删除门店库
func TaobaoPlaceStoregroupDelete(clt *core.SDKClient, req *store.TaobaoPlaceStoregroupDeleteAPIRequest, session string) (*store.TaobaoPlaceStoregroupDeleteAPIResponse, error) {
	var resp store.TaobaoPlaceStoregroupDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
