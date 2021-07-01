package tttm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmProduceSyncAPIResponse
天天特卖生产进度同步 API返回值
aliyun.industry.tttm.produce.sync

天天特卖生产进度同步 */
type AliyunIndustryTttmProduceSyncAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmProduceSyncAPIResponseModel
}

// AliyunIndustryTttmProduceSyncAPIResponseModel is 天天特卖生产进度同步 成功返回结果
type AliyunIndustryTttmProduceSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_produce_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 状态
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
