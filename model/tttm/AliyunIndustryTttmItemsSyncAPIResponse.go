package tttm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunIndustryTttmItemsSyncAPIResponse 天天特卖商品信息同步 API返回值
// aliyun.industry.tttm.items.sync
//
// 天天特卖商品信息同步
type AliyunIndustryTttmItemsSyncAPIResponse struct {
	model.CommonResponse
	AliyunIndustryTttmItemsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunIndustryTttmItemsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunIndustryTttmItemsSyncAPIResponseModel).Reset()
}

// AliyunIndustryTttmItemsSyncAPIResponseModel is 天天特卖商品信息同步 成功返回结果
type AliyunIndustryTttmItemsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_industry_tttm_items_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功失败标识
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliyunIndustryTttmItemsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAliyunIndustryTttmItemsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunIndustryTttmItemsSyncAPIResponse)
	},
}

// GetAliyunIndustryTttmItemsSyncAPIResponse 从 sync.Pool 获取 AliyunIndustryTttmItemsSyncAPIResponse
func GetAliyunIndustryTttmItemsSyncAPIResponse() *AliyunIndustryTttmItemsSyncAPIResponse {
	return poolAliyunIndustryTttmItemsSyncAPIResponse.Get().(*AliyunIndustryTttmItemsSyncAPIResponse)
}

// ReleaseAliyunIndustryTttmItemsSyncAPIResponse 将 AliyunIndustryTttmItemsSyncAPIResponse 保存到 sync.Pool
func ReleaseAliyunIndustryTttmItemsSyncAPIResponse(v *AliyunIndustryTttmItemsSyncAPIResponse) {
	v.Reset()
	poolAliyunIndustryTttmItemsSyncAPIResponse.Put(v)
}
