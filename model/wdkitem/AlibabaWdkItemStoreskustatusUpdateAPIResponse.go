package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskustatusUpdateAPIResponse 修改门店商品状态 API返回值
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
type AlibabaWdkItemStoreskustatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemStoreskustatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskustatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemStoreskustatusUpdateAPIResponseModel).Reset()
}

// AlibabaWdkItemStoreskustatusUpdateAPIResponseModel is 修改门店商品状态 成功返回结果
type AlibabaWdkItemStoreskustatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storeskustatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemStoreskustatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskustatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemStoreskustatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskustatusUpdateAPIResponse)
	},
}

// GetAlibabaWdkItemStoreskustatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemStoreskustatusUpdateAPIResponse
func GetAlibabaWdkItemStoreskustatusUpdateAPIResponse() *AlibabaWdkItemStoreskustatusUpdateAPIResponse {
	return poolAlibabaWdkItemStoreskustatusUpdateAPIResponse.Get().(*AlibabaWdkItemStoreskustatusUpdateAPIResponse)
}

// ReleaseAlibabaWdkItemStoreskustatusUpdateAPIResponse 将 AlibabaWdkItemStoreskustatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemStoreskustatusUpdateAPIResponse(v *AlibabaWdkItemStoreskustatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemStoreskustatusUpdateAPIResponse.Put(v)
}
