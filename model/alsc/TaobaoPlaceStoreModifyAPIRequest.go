package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreModifyAPIRequest
商家修改线下门店 API请求
taobao.place.store.modify

用于商家修改线下门店信息 */
type TaobaoPlaceStoreModifyAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeUpdate *StoreUpdateTopDto
}

// New
