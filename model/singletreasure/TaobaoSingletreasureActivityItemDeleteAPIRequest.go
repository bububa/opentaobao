package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityItemDeleteAPIRequest
删除单品优惠接口 API请求
taobao.singletreasure.activity.item.delete

删除单品优惠接口 */
type TaobaoSingletreasureActivityItemDeleteAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 商品Id
	_itemId int64
}

// New
