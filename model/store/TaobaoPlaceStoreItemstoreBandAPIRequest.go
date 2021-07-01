package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreItemstoreBandAPIRequest
门店商品关联绑定接口 API请求
taobao.place.store.itemstore.band

商品和多个门店关系绑定接口 */
type TaobaoPlaceStoreItemstoreBandAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
	// 门店id
	_storeIds []int64
	// 操作类型
	_actionType string
}

// New
