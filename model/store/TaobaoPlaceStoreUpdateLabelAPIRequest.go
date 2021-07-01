package store

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPlaceStoreUpdateLabelAPIRequest
商户门店标签更新接口 API请求
taobao.place.store.update.label

更新商户门店标签（服务、权益、标签）接口 */
type TaobaoPlaceStoreUpdateLabelAPIRequest struct {
	model.Params
	// 门店id
	_storeId int64
	// 标签id
	_labelIdList []int64
	// 行业code
	_businessCode string
	// 标签类型
	_labelType string
}

// New
