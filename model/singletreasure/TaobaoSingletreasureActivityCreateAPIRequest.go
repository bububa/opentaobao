package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityCreateAPIRequest
活动创建接口 API请求
taobao.singletreasure.activity.create

创建优惠活动 */
type TaobaoSingletreasureActivityCreateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// New
