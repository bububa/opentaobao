package tttm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmPlanSyncAPIResponse
天天特卖生产计划单同步 API返回值
aliyun.industry.tttm.plan.sync

ERP系统向天天特卖同步生产计划单的数据 */
type AliyunIndustryTttmPlanSyncAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmPlanSyncAPIResponseModel
}

// AliyunIndustryTttmPlanSyncAPIResponseModel is 天天特卖生产计划单同步 成功返回结果
type AliyunIndustryTttmPlanSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_plan_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
