package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest
加密招商一网能支付入参 API请求
alibaba.damai.maitix.distribution.cmb.paramencrypt

encryptParam4Cmb */
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest struct {
	model.Params
	// 入参param
	_param *DisEncrypt4CmbParam
}

// New
