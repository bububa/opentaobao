package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApprovalUpdateAPIRequest
更新审批单 API请求
alitrip.btrip.approval.update

更新审批单 */
type AlitripBtripApprovalUpdateAPIRequest struct {
	model.Params
	// 审批请求对象
	_approveApplyRequest *OpenApproveApplyRq
}

// New
