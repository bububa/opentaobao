package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixseattokenquery 分销商选座获取qtoken
// alibaba.damai.maitix.seat.token.query
//
// 选座分销，分销商查询token
func Alibabadamaimaitixseattokenquery(clt *core.SDKClient, req *maitix.AlibabadamaimaitixseattokenqueryAPIRequest, session string) (*maitix.AlibabadamaimaitixseattokenqueryAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixseattokenqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
