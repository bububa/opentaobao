package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolActivityCreateAPIResponse 创建活动新接口 API返回值
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabaHmMarketingItempoolActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolActivityCreateAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolActivityCreateAPIResponseModel is 创建活动新接口 成功返回结果
type AlibabaHmMarketingItempoolActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_activity_create_response"`
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
func (m *AlibabaHmMarketingItempoolActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.FailCode = ""
	m.Data = 0
	m.IsSuccess = false
}

var poolAlibabaHmMarketingItempoolActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolActivityCreateAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolActivityCreateAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolActivityCreateAPIResponse
func GetAlibabaHmMarketingItempoolActivityCreateAPIResponse() *AlibabaHmMarketingItempoolActivityCreateAPIResponse {
	return poolAlibabaHmMarketingItempoolActivityCreateAPIResponse.Get().(*AlibabaHmMarketingItempoolActivityCreateAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolActivityCreateAPIResponse 将 AlibabaHmMarketingItempoolActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolActivityCreateAPIResponse(v *AlibabaHmMarketingItempoolActivityCreateAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolActivityCreateAPIResponse.Put(v)
}
