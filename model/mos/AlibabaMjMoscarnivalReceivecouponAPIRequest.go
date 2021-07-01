package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMoscarnivalReceivecouponAPIRequest
根据手机号码领券 API请求
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券 */
type AlibabaMjMoscarnivalReceivecouponAPIRequest struct {
	model.Params
	// 手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// New
