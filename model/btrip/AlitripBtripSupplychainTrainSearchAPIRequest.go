package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainTrainSearchAPIRequest
【商旅】火车票订单查询 API请求
alitrip.btrip.supplychain.train.search

【商旅】火车票订单查询 */
type AlitripBtripSupplychainTrainSearchAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiSearchRq
}

// New
