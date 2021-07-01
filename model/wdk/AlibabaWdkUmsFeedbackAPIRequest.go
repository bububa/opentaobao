package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkUmsFeedbackAPIRequest
质量反馈（入库辅助）-ERP下发单 API请求
alibaba.wdk.ums.feedback

质量反馈（入库辅助）-ERP下发单 */
type AlibabaWdkUmsFeedbackAPIRequest struct {
	model.Params
	// 质量反馈请求dto
	_erpFeedbackdto *ErpFeedbackDto
}

// New
