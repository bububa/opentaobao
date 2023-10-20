package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingOpenVersionApplyAPIResponse 数据同步版本号申请 API返回值
// alibaba.wdk.marketing.open.version.apply
//
// 数据同步版本号申请
type AlibabaWdkMarketingOpenVersionApplyAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenVersionApplyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenVersionApplyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingOpenVersionApplyAPIResponseModel).Reset()
}

// AlibabaWdkMarketingOpenVersionApplyAPIResponseModel is 数据同步版本号申请 成功返回结果
type AlibabaWdkMarketingOpenVersionApplyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_version_apply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 版本号申请结果
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingOpenVersionApplyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingOpenVersionApplyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingOpenVersionApplyAPIResponse)
	},
}

// GetAlibabaWdkMarketingOpenVersionApplyAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingOpenVersionApplyAPIResponse
func GetAlibabaWdkMarketingOpenVersionApplyAPIResponse() *AlibabaWdkMarketingOpenVersionApplyAPIResponse {
	return poolAlibabaWdkMarketingOpenVersionApplyAPIResponse.Get().(*AlibabaWdkMarketingOpenVersionApplyAPIResponse)
}

// ReleaseAlibabaWdkMarketingOpenVersionApplyAPIResponse 将 AlibabaWdkMarketingOpenVersionApplyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingOpenVersionApplyAPIResponse(v *AlibabaWdkMarketingOpenVersionApplyAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingOpenVersionApplyAPIResponse.Put(v)
}
