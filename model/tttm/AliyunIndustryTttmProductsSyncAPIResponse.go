package tttm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliyunIndustryTttmProductsSyncAPIResponse
天天特卖货品信息同步 API返回值
aliyun.industry.tttm.products.sync

天天特卖货品信息同步 */
type AliyunIndustryTttmProductsSyncAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmProductsSyncAPIResponseModel
}

// AliyunIndustryTttmProductsSyncAPIResponseModel is 天天特卖货品信息同步 成功返回结果
type AliyunIndustryTttmProductsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_products_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功失败标识
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
