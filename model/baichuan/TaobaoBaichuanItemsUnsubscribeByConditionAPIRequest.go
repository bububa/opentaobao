package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest
根据条件删除订阅关系 API请求
taobao.baichuan.items.unsubscribe.by.condition

根据条件删除订阅关系 */
type TaobaoBaichuanItemsUnsubscribeByConditionAPIRequest struct {
	model.Params
	// 删除条件
	_condition *Condition
}

// New
