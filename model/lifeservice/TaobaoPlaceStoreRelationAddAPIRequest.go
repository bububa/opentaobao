package lifeservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreRelationAddAPIRequest
门店关系新增 API请求
taobao.place.store.relation.add

新增授权用户的门店关系信息 */
type TaobaoPlaceStoreRelationAddAPIRequest struct {
	model.Params
	// 关系B的门店ID
	_outerId string
	// 关系类型(outer_id是主, store_id是从)
	_relationType int64
	// 门店ID
	_storeId int64
}

// New
