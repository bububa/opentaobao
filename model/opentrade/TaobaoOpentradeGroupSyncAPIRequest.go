package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeGroupSyncAPIRequest
组团购场景创建或更新组团活动 API请求
taobao.opentrade.group.sync

组团购场景中创建团购活动 */
type TaobaoOpentradeGroupSyncAPIRequest struct {
	model.Params
	// 组团活动开始时间
	_startTime string
	// 组团活动结束时间
	_endTime string
	// 成团有效期，单位为妙
	_expiration int64
	// 成团的目标人数
	_goal int64
	// 组团类型，0：拼团；1：团购
	_groupType int64
	// 是否任何账号可开团。whitelist：仅白名单账号可开团  all：任何账号可开团
	_allowType string
	// 允许开团的淘宝账号列表
	_allowWhiteList []string
	// 组团类型为团购，可限制团长针对一个商品的开团数量上限
	_openLimit int64
	// 未成团处理办法，close：系统关单；continue：订单继续下行
	_failProcess string
	// 组团购买的折扣价，单位为分
	_discountPrice int64
	// 商品ID
	_itemId int64
	// 组团活动id
	_groupActivityId int64
}

// New
