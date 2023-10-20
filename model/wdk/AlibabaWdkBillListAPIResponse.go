package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBillListAPIResponse 五道口账单拉取接口 API返回值
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
type AlibabaWdkBillListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBillListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBillListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBillListAPIResponseModel).Reset()
}

// AlibabaWdkBillListAPIResponseModel is 五道口账单拉取接口 成功返回结果
type AlibabaWdkBillListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bill_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应结果
	ApiResult *AlibabaWdkBillListApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBillListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkBillListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBillListAPIResponse)
	},
}

// GetAlibabaWdkBillListAPIResponse 从 sync.Pool 获取 AlibabaWdkBillListAPIResponse
func GetAlibabaWdkBillListAPIResponse() *AlibabaWdkBillListAPIResponse {
	return poolAlibabaWdkBillListAPIResponse.Get().(*AlibabaWdkBillListAPIResponse)
}

// ReleaseAlibabaWdkBillListAPIResponse 将 AlibabaWdkBillListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBillListAPIResponse(v *AlibabaWdkBillListAPIResponse) {
	v.Reset()
	poolAlibabaWdkBillListAPIResponse.Put(v)
}
