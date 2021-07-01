package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemSubscribeAPIRequest
单个商品订阅 API请求
taobao.baichuan.item.subscribe

百川单个商品订阅 */
type TaobaoBaichuanItemSubscribeAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
