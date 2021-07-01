package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeActivitySyncAPIRequest
尖货交易活动信息同步 API请求
taobao.opentrade.activity.sync

尖货交易活动信息配置，创建或更新活动信息
在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买； */
type TaobaoOpentradeActivitySyncAPIRequest struct {
	model.Params
	// 排队活动ID，排队时如传入，这里需要填写；若未传，这里也可以不传
	_activityId string
	// 活动开始时间
	_startTime string
	// 活动结束时间（全流程结束时间，非排队结束时间）
	_endTime string
	// 活动名称
	_activityName string
	// 活动关联的商品列表，使用逗号(,)分割
	_itemIdList []int64
}

// New
