package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripApprovalModify
修改审批单
alitrip.btrip.approval.modify

修改审批单 */
func AlitripBtripApprovalModify(clt *core.SDKClient, req *btrip.AlitripBtripApprovalModifyAPIRequest, session string) (*btrip.AlitripBtripApprovalModifyAPIResponse, error) {
	var resp btrip.AlitripBtripApprovalModifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
