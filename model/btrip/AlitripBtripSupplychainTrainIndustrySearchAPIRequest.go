package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripSupplychainTrainIndustrySearchAPIRequest
火车票行业搜索接口 API请求
alitrip.btrip.supplychain.train.industry.search

【商旅】火车票行业搜索接口 */
type AlitripBtripSupplychainTrainIndustrySearchAPIRequest struct {
	model.Params
	// 入参
	_rq *TrainSearchRq
}

// New
