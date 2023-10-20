package tttm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tttm"
)

// AliyunIndustryTttmPlanSync 天天特卖生产计划单同步
// aliyun.industry.tttm.plan.sync
//
// ERP系统向天天特卖同步生产计划单的数据
func AliyunIndustryTttmPlanSync(clt *core.SDKClient, req *tttm.AliyunIndustryTttmPlanSyncAPIRequest, resp *tttm.AliyunIndustryTttmPlanSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
