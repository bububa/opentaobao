package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripApprovalNewAPIRequest
新建审批单 API请求
alitrip.btrip.approval.new

用户新建审批单 */
type AlitripBtripApprovalNewAPIRequest struct {
	model.Params
	// 申请单
	_addApplyRequest *OpenAddApplyRq
}

// New
