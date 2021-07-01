package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStorerelatesubAddAPIRequest
门店和子门店关系新增 API请求
taobao.place.storerelatesub.add

门店和子门店关系新增 */
type TaobaoPlaceStorerelatesubAddAPIRequest struct {
	model.Params
	// 门店Id
	_storeId int64
	// 子门店Id
	_subStoreIds []int64
}

// New
