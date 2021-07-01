package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSeakingTaskReportAPIRequest
跳转任务发布成功商品ID回传 API请求
alibaba.seaking.task.report

跳转任务发布成功商品ID回传 */
type AlibabaSeakingTaskReportAPIRequest struct {
	model.Params
	// 上报数据详情
	_reportDetail []TaskDetailReportDto
	// 任务类型(title/image)
	_taskType string
	// 用户token
	_token string
	// token来源站点
	_tokenFrom string
}

// New
