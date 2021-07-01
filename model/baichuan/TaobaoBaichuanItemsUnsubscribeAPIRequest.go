package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemsUnsubscribeAPIRequest
批量删除商品订阅 API请求
taobao.baichuan.items.unsubscribe

批量删除商品订阅 */
type TaobaoBaichuanItemsUnsubscribeAPIRequest struct {
	model.Params
	// 删除的商品id
	_itemIds []int64
}

// New
