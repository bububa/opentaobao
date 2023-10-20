package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolActivityCreateAPIResponse 创建活动新接口 API返回值
// alibaba.wdk.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabaWdkMarketingItempoolActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolActivityCreateAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolActivityCreateAPIResponseModel is 创建活动新接口 成功返回结果
type AlibabaWdkMarketingItempoolActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// data
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.FailCode = ""
	m.Data = 0
	m.IsSuccess = false
}

var poolAlibabaWdkMarketingItempoolActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolActivityCreateAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolActivityCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolActivityCreateAPIResponse
func GetAlibabaWdkMarketingItempoolActivityCreateAPIResponse() *AlibabaWdkMarketingItempoolActivityCreateAPIResponse {
	return poolAlibabaWdkMarketingItempoolActivityCreateAPIResponse.Get().(*AlibabaWdkMarketingItempoolActivityCreateAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolActivityCreateAPIResponse 将 AlibabaWdkMarketingItempoolActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolActivityCreateAPIResponse(v *AlibabaWdkMarketingItempoolActivityCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolActivityCreateAPIResponse.Put(v)
}
