package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCityCarApplyApproveAPIRequest
三方市内用车申请单审批 API请求
alitrip.btrip.city.car.apply.approve

三方市内用车申请单审批 */
type AlitripBtripCityCarApplyApproveAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyApproveRq
}

// New
