package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// TaobaoPlaceStoreRelationQuery 门店关系查询
// taobao.place.store.relation.query
//
// 查询门店关系
func TaobaoPlaceStoreRelationQuery(clt *core.SDKClient, req *store.TaobaoPlaceStoreRelationQueryAPIRequest, resp *store.TaobaoPlaceStoreRelationQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
