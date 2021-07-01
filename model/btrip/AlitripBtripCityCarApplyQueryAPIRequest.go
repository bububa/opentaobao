package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCityCarApplyQueryAPIRequest
三方市内用车申请单查询 API请求
alitrip.btrip.city.car.apply.query

三方市内用车申请单查询 */
type AlitripBtripCityCarApplyQueryAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyQueryRq
}

// New
