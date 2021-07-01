package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWestcrmJobStatusUpdateAPIRequest
更新工单状态 API请求
alibaba.westcrm.job.status.update

更新工单处理状态 */
type AlibabaWestcrmJobStatusUpdateAPIRequest struct {
	model.Params
	// 状态
	_status int64
	// 园区id
	_campusId int64
	// 反馈id
	_crmComplaintId int64
	// 回复内容
	_replyContent string
}

// New
