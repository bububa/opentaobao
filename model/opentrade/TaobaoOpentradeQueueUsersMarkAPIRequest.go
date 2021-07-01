package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeQueueUsersMarkAPIRequest
尖货交易可购买用户标记 API请求
taobao.opentrade.queue.users.mark

尖货交易用户标记信息回传，根据openId标记用户可购买商品 */
type TaobaoOpentradeQueueUsersMarkAPIRequest struct {
	model.Params
	// 用户状态，可任意传入，后续查询返回
	_status string
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 排队商品SKU ID，不存在传0
	_skuId int64
	// 排队商品ID
	_itemId int64
	// 本次待标记的用户列表，多个以逗号(,)分割，最大20个
	_openUserIds []string
	// 是否目标用户，传入true后，用户可购买商品
	_hit bool
}

// New
