package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemsSubscribeAPIRequest
百川批量商品订阅 API请求
taobao.baichuan.items.subscribe

百川批量添加订阅的商品 */
type TaobaoBaichuanItemsSubscribeAPIRequest struct {
	model.Params
	// 订阅的商品id列表
	_itemIds []int64
}

// New
