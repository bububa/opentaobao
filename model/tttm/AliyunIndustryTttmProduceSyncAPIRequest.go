package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmProduceSyncAPIRequest
天天特卖生产进度同步 API请求
aliyun.industry.tttm.produce.sync

天天特卖生产进度同步 */
type AliyunIndustryTttmProduceSyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// New
