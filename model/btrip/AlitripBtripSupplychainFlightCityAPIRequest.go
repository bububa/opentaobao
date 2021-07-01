package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainFlightCityAPIRequest
机场数据查询 API请求
alitrip.btrip.supplychain.flight.city

机场数据查询 */
type AlitripBtripSupplychainFlightCityAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSuggestRq
}

// New
