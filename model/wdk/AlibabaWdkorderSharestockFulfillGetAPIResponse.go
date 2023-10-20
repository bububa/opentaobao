package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockFulfillGetAPIResponse 商户订单履约数据获取 API返回值
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
type AlibabaWdkorderSharestockFulfillGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockFulfillGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockFulfillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockFulfillGetAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockFulfillGetAPIResponseModel is 商户订单履约数据获取 成功返回结果
type AlibabaWdkorderSharestockFulfillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_fulfill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *MaochaoOrderFulfillQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockFulfillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockFulfillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockFulfillGetAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockFulfillGetAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockFulfillGetAPIResponse
func GetAlibabaWdkorderSharestockFulfillGetAPIResponse() *AlibabaWdkorderSharestockFulfillGetAPIResponse {
	return poolAlibabaWdkorderSharestockFulfillGetAPIResponse.Get().(*AlibabaWdkorderSharestockFulfillGetAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockFulfillGetAPIResponse 将 AlibabaWdkorderSharestockFulfillGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockFulfillGetAPIResponse(v *AlibabaWdkorderSharestockFulfillGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockFulfillGetAPIResponse.Put(v)
}
