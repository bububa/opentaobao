package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixdistributioncmbparamencrypt 加密招商一网能支付入参
// alibaba.damai.maitix.distribution.cmb.paramencrypt
//
// encryptParam4Cmb
func Alibabadamaimaitixdistributioncmbparamencrypt(clt *core.SDKClient, req *maitix.AlibabadamaimaitixdistributioncmbparamencryptAPIRequest, session string) (*maitix.AlibabadamaimaitixdistributioncmbparamencryptAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixdistributioncmbparamencryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
