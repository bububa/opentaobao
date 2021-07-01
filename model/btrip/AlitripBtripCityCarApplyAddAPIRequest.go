package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCityCarApplyAddAPIRequest
三方市内用车申请单同步 API请求
alitrip.btrip.city.car.apply.add

三方市内用车申请单同步 */
type AlitripBtripCityCarApplyAddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *CityCarApplyAddRq
}

// New
