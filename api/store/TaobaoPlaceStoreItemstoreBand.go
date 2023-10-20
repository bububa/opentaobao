package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoreitemstoreband 门店商品关联绑定接口
// taobao.place.store.itemstore.band
//
// 商品和多个门店关系绑定接口
func Taobaoplacestoreitemstoreband(clt *core.SDKClient, req *store.TaobaoplacestoreitemstorebandAPIRequest, session string) (*store.TaobaoplacestoreitemstorebandAPIResponse, error) {
	var resp store.TaobaoplacestoreitemstorebandAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
