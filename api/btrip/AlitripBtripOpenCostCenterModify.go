package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcentermodify 修改成本中心
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
func Alitripbtripopencostcentermodify(clt *core.SDKClient, req *btrip.AlitripbtripopencostcentermodifyAPIRequest, session string) (*btrip.AlitripbtripopencostcentermodifyAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcentermodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
