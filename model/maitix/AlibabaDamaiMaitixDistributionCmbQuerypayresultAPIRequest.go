package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest
查询招行支付状态api API请求
alibaba.damai.maitix.distribution.cmb.querypayresult

queryPayResult */
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest struct {
	model.Params
	// 入参param
	_param *QueryPayResultParam
}

// New
