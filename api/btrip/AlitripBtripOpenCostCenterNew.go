package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenternew 新增成本中心
// alitrip.btrip.open.cost.center.new
//
// 新增成本中心
func Alitripbtripopencostcenternew(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenternewAPIRequest, session string) (*btrip.AlitripbtripopencostcenternewAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenternewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
