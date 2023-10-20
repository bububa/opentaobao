package tttm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmProductsSyncAPIResponse 天天特卖货品信息同步 API返回值
// aliyun.industry.tttm.products.sync
//
// 天天特卖货品信息同步
type AliyunIndustryTttmProductsSyncAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmProductsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunIndustryTttmProductsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunIndustryTttmProductsSyncAPIResponseModel).Reset()
}

// AliyunIndustryTttmProductsSyncAPIResponseModel is 天天特卖货品信息同步 成功返回结果
type AliyunIndustryTttmProductsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_products_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功失败标识
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliyunIndustryTttmProductsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAliyunIndustryTttmProductsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunIndustryTttmProductsSyncAPIResponse)
	},
}

// GetAliyunIndustryTttmProductsSyncAPIResponse 从 sync.Pool 获取 AliyunIndustryTttmProductsSyncAPIResponse
func GetAliyunIndustryTttmProductsSyncAPIResponse() *AliyunIndustryTttmProductsSyncAPIResponse {
	return poolAliyunIndustryTttmProductsSyncAPIResponse.Get().(*AliyunIndustryTttmProductsSyncAPIResponse)
}

// ReleaseAliyunIndustryTttmProductsSyncAPIResponse 将 AliyunIndustryTttmProductsSyncAPIResponse 保存到 sync.Pool
func ReleaseAliyunIndustryTttmProductsSyncAPIResponse(v *AliyunIndustryTttmProductsSyncAPIResponse) {
	v.Reset()
	poolAliyunIndustryTttmProductsSyncAPIResponse.Put(v)
}
