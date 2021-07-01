package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApprovalModifyAPIRequest
修改审批单 API请求
alitrip.btrip.approval.modify

修改审批单 */
type AlitripBtripApprovalModifyAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenApiNewApplyRq
}

// New
