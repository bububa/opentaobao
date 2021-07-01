package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMjMoscarnivalReceiveencryptAPIRequest
根据加密手机号领券 API请求
alibaba.mj.moscarnival.receiveencrypt

根据加密手机号领券 */
type AlibabaMjMoscarnivalReceiveencryptAPIRequest struct {
	model.Params
	// 加密手机号码
	_mobile string
	// 活动id
	_activityId int64
}

// New
