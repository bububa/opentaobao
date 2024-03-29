package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillBoxPostBackBoxAPIResponse RT收箱回传 API返回值
// alibaba.wdk.fulfill.box.post.back.box
//
// RT收箱后，信息同步履约，履约同通知UMS 容器管理
type AlibabaWdkFulfillBoxPostBackBoxAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillBoxPostBackBoxAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBoxPostBackBoxAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillBoxPostBackBoxAPIResponseModel).Reset()
}

// AlibabaWdkFulfillBoxPostBackBoxAPIResponseModel is RT收箱回传 成功返回结果
type AlibabaWdkFulfillBoxPostBackBoxAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_box_post_back_box_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// fulfillLogisticSingleResult
	FulfillLogisticSingleResult *FulfillLogisticDefaultResult `json:"fulfill_logistic_single_result,omitempty" xml:"fulfill_logistic_single_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillBoxPostBackBoxAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FulfillLogisticSingleResult = nil
}

var poolAlibabaWdkFulfillBoxPostBackBoxAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillBoxPostBackBoxAPIResponse)
	},
}

// GetAlibabaWdkFulfillBoxPostBackBoxAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillBoxPostBackBoxAPIResponse
func GetAlibabaWdkFulfillBoxPostBackBoxAPIResponse() *AlibabaWdkFulfillBoxPostBackBoxAPIResponse {
	return poolAlibabaWdkFulfillBoxPostBackBoxAPIResponse.Get().(*AlibabaWdkFulfillBoxPostBackBoxAPIResponse)
}

// ReleaseAlibabaWdkFulfillBoxPostBackBoxAPIResponse 将 AlibabaWdkFulfillBoxPostBackBoxAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillBoxPostBackBoxAPIResponse(v *AlibabaWdkFulfillBoxPostBackBoxAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillBoxPostBackBoxAPIResponse.Put(v)
}
