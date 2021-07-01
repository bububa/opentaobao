package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripOpenCostCenterModify
修改成本中心
alitrip.btrip.open.cost.center.modify

修改成本中心 */
func AlitripBtripOpenCostCenterModify(clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterModifyAPIRequest, session string) (*btrip.AlitripBtripOpenCostCenterModifyAPIResponse, error) {
	var resp btrip.AlitripBtripOpenCostCenterModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
