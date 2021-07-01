package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityDeleteAPIRequest
删除活动接口 API请求
taobao.singletreasure.activity.delete

删除优惠活动 */
type TaobaoSingletreasureActivityDeleteAPIRequest struct {
	model.Params
	// 活动Id
	_activityId int64
}

// New
