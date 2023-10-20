package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtipcostcenterquery 查询外部成本中心
// alitrip.btip.cost.center.query
//
// 查询外部成本中心
func Alitripbtipcostcenterquery(clt *core.SDKClient, req *btrip.AlitripbtipcostcenterqueryAPIRequest, session string) (*btrip.AlitripbtipcostcenterqueryAPIResponse, error) {
	var resp btrip.AlitripbtipcostcenterqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
