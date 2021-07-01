package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreDeleteAPIRequest
线下门店删除 API请求
taobao.place.store.delete

用于商家删除线下门店 */
type TaobaoPlaceStoreDeleteAPIRequest struct {
	model.Params
	// 门店id
	_storeId int64
}

// New
