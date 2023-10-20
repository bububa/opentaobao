package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionCmbParamencrypt 加密招商一网能支付入参
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
func AlibabaDamaiMaitixDistributionCmbParamencrypt(clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionCmbParamencryptAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
