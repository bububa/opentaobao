package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixorderdistributioncreate 大麦-新分销下单
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
func Alibabadamaimaitixorderdistributioncreate(clt *core.SDKClient, req *maitix.AlibabadamaimaitixorderdistributioncreateAPIRequest, session string) (*maitix.AlibabadamaimaitixorderdistributioncreateAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixorderdistributioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
