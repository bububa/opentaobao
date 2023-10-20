package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcentertransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func Alitripbtripopencostcentertransfer(clt *core.SDKClient, req *btrip.AlitripbtripopencostcentertransferAPIRequest, session string) (*btrip.AlitripbtripopencostcentertransferAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcentertransferAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
