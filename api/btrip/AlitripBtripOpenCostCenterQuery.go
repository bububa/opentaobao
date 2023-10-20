package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenterquery 查询成本中心
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
func Alitripbtripopencostcenterquery(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenterqueryAPIRequest, session string) (*btrip.AlitripbtripopencostcenterqueryAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenterqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
