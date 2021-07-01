package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleParttimeJobsyncAPIRequest
兼职岗位同步 API请求
alibaba.idle.parttime.jobsync

服务商同步岗位信息给闲鱼 */
type AlibabaIdleParttimeJobsyncAPIRequest struct {
	model.Params
	// 岗位列表
	_jobList []PartTimeJob
	// 同步数据的时间
	_syncTime int64
}

// New
