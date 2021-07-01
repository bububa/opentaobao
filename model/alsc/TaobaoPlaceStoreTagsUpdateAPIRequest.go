package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreTagsUpdateAPIRequest
门店打标去标 API请求
taobao.place.store.tags.update

门店打标去标 */
type TaobaoPlaceStoreTagsUpdateAPIRequest struct {
	model.Params
	// 门店信息
	_storeUpdate *StoreUpdateTopDto
	// 新增标list
	_addTags []int64
	// 删除标list
	_removeTags []int64
}

// New
