package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreRelationQueryAPIRequest
门店关系查询 API请求
taobao.place.store.relation.query

查询门店关系 */
type TaobaoPlaceStoreRelationQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramStoreRelationSimpleQuery *StoreRelationSimpleQuery
}

// New
