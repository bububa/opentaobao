package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStorerelatesubDeleteAPIRequest
门店和子门店关系删除 API请求
taobao.place.storerelatesub.delete

门店和子门店关系删除 */
type TaobaoPlaceStorerelatesubDeleteAPIRequest struct {
	model.Params
	// 门店Id
	_storeId int64
	// 子门店id
	_subStoreIds []int64
}

// New
