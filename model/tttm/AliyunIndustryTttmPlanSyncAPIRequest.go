package tttm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmPlanSyncAPIRequest
天天特卖生产计划单同步 API请求
aliyun.industry.tttm.plan.sync

ERP系统向天天特卖同步生产计划单的数据 */
type AliyunIndustryTttmPlanSyncAPIRequest struct {
	model.Params
	// 计划单
	_syncPlan *SyncPlanDto
}

// New
