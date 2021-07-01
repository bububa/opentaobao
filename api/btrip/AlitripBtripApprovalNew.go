package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripApprovalNew
新建审批单
alitrip.btrip.approval.new

用户新建审批单 */
func AlitripBtripApprovalNew(clt *core.SDKClient, req *btrip.AlitripBtripApprovalNewAPIRequest, session string) (*btrip.AlitripBtripApprovalNewAPIResponse, error) {
	var resp btrip.AlitripBtripApprovalNewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
