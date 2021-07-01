package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainTrainCityAPIRequest
火车站数据查询 API请求
alitrip.btrip.supplychain.train.city

火车站数据查询 */
type AlitripBtripSupplychainTrainCityAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenSuggestRq
}

// New
