package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixeticketdistributionquery 分销电子票查询接口
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
func Alibabadamaimaitixeticketdistributionquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixeticketdistributionqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixeticketdistributionqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixeticketdistributionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
