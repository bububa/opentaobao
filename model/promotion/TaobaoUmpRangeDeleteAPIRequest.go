package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpRangeDeleteAPIRequest
删除活动范围 API请求
taobao.ump.range.delete

去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置 */
type TaobaoUmpRangeDeleteAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
	// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
	_type int64
	// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
	_ids string
}

// New
