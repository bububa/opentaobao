package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemUnsubscribeAPIRequest
单个删除订阅关系 API请求
taobao.baichuan.item.unsubscribe

删除单个商品订阅关系 */
type TaobaoBaichuanItemUnsubscribeAPIRequest struct {
	model.Params
	// 商品id
	_itemId int64
}

// New
