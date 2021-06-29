package tttm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天天特卖生产计划单同步 APIResponse
aliyun.industry.tttm.plan.sync

ERP系统向天天特卖同步生产计划单的数据
*/
type AliyunIndustryTttmPlanSyncAPIResponse struct {
    model.CommonResponse
    AliyunIndustryTttmPlanSyncResponse
}

type AliyunIndustryTttmPlanSyncResponse struct {
    XMLName xml.Name `xml:"aliyun_industry_tttm_plan_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 状态
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
