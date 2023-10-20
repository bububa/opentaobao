package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkStockPublishAPIResponse 五道口库存发布接口（针对外部渠道） API返回值
// alibaba.wdk.stock.publish
//
// 五道口库存发布接口（针对外部渠道）
type AlibabaWdkStockPublishAPIResponse struct {
	model.CommonResponse
	AlibabaWdkStockPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkStockPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkStockPublishAPIResponseModel).Reset()
}

// AlibabaWdkStockPublishAPIResponseModel is 五道口库存发布接口（针对外部渠道） 成功返回结果
type AlibabaWdkStockPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_stock_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// errorMsg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkStockPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.Message = ""
	m.IsSuccess = false
}

var poolAlibabaWdkStockPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkStockPublishAPIResponse)
	},
}

// GetAlibabaWdkStockPublishAPIResponse 从 sync.Pool 获取 AlibabaWdkStockPublishAPIResponse
func GetAlibabaWdkStockPublishAPIResponse() *AlibabaWdkStockPublishAPIResponse {
	return poolAlibabaWdkStockPublishAPIResponse.Get().(*AlibabaWdkStockPublishAPIResponse)
}

// ReleaseAlibabaWdkStockPublishAPIResponse 将 AlibabaWdkStockPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkStockPublishAPIResponse(v *AlibabaWdkStockPublishAPIResponse) {
	v.Reset()
	poolAlibabaWdkStockPublishAPIResponse.Put(v)
}
