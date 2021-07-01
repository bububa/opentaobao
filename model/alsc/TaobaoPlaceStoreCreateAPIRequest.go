package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreCreateAPIRequest
商户门店创建接口 API请求
taobao.place.store.create

用于商家创建线下门店 */
type TaobaoPlaceStoreCreateAPIRequest struct {
	model.Params
	// 门店创建入参
	_storeCreate *StoreUpdateTopDto
}

// New
