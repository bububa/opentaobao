package tttm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunindustrytttmplansyncAPIResponse 天天特卖生产计划单同步 API返回值
// aliyun.industry.tttm.plan.sync
//
// ERP系统向天天特卖同步生产计划单的数据
type AliyunindustrytttmplansyncAPIResponse struct {
	model.CommonResponse
	AliyunindustrytttmplansyncAPIResponseModel
}

// AliyunindustrytttmplansyncAPIResponseModel is 天天特卖生产计划单同步 成功返回结果
type AliyunindustrytttmplansyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_plan_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
