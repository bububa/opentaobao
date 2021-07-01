package singletreasure

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSingletreasureActivityUpdateAPIRequest
修改活动接口 API请求
taobao.singletreasure.activity.update

修改活动接口 */
type TaobaoSingletreasureActivityUpdateAPIRequest struct {
	model.Params
	// 系统入参
	_activityInfo *ActivityInfoCreateDto
}

// New
