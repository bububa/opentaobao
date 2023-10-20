package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixdistributiondeliverycalculate 计算渠道用户下单快递费
// alibaba.damai.maitix.distribution.delivery.calculate
//
// 计算渠道用户下单快递费
func Alibabadamaimaitixdistributiondeliverycalculate(clt *core.SDKClient, req *maitix.AlibabadamaimaitixdistributiondeliverycalculateAPIRequest, session string) (*maitix.AlibabadamaimaitixdistributiondeliverycalculateAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixdistributiondeliverycalculateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
