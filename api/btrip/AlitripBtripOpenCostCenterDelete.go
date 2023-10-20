package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripopencostcenterdelete 删除成本中心
// alitrip.btrip.open.cost.center.delete
//
// 删除成本中心
func Alitripbtripopencostcenterdelete(clt *core.SDKClient, req *btrip.AlitripbtripopencostcenterdeleteAPIRequest, session string) (*btrip.AlitripbtripopencostcenterdeleteAPIResponse, error) {
	var resp btrip.AlitripbtripopencostcenterdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
