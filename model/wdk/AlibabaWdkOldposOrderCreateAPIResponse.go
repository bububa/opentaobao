package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOldposOrderCreateAPIResponse 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 API返回值
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
type AlibabaWdkOldposOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkOldposOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkOldposOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkOldposOrderCreateAPIResponseModel).Reset()
}

// AlibabaWdkOldposOrderCreateAPIResponseModel is 淘鲜达外部商户老pos机产生的订单同步进淘鲜达 成功返回结果
type AlibabaWdkOldposOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_oldpos_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *PosOrderCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkOldposOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkOldposOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOldposOrderCreateAPIResponse)
	},
}

// GetAlibabaWdkOldposOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkOldposOrderCreateAPIResponse
func GetAlibabaWdkOldposOrderCreateAPIResponse() *AlibabaWdkOldposOrderCreateAPIResponse {
	return poolAlibabaWdkOldposOrderCreateAPIResponse.Get().(*AlibabaWdkOldposOrderCreateAPIResponse)
}

// ReleaseAlibabaWdkOldposOrderCreateAPIResponse 将 AlibabaWdkOldposOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkOldposOrderCreateAPIResponse(v *AlibabaWdkOldposOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkOldposOrderCreateAPIResponse.Put(v)
}
