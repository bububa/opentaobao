package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixdistributioncmbquerypayresult 查询招行支付状态api
// alibaba.damai.maitix.distribution.cmb.querypayresult
//
// queryPayResult
func Alibabadamaimaitixdistributioncmbquerypayresult(clt *core.SDKClient, req *maitix.AlibabadamaimaitixdistributioncmbquerypayresultAPIRequest, session string) (*maitix.AlibabadamaimaitixdistributioncmbquerypayresultAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixdistributioncmbquerypayresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
