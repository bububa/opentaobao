package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripCostCenterGet
获取用户费用归属
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表 */
func AlitripBtripCostCenterGet(clt *core.SDKClient, req *btrip.AlitripBtripCostCenterGetAPIRequest, session string) (*btrip.AlitripBtripCostCenterGetAPIResponse, error) {
	var resp btrip.AlitripBtripCostCenterGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
