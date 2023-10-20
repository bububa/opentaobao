package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskuUpdateAPIResponse 五道口商品中心门店商品修改 API返回值
// alibaba.wdk.item.storesku.update
//
// 五道口商品中心门店商品修改
type AlibabaWdkItemStoreskuUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemStoreskuUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskuUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemStoreskuUpdateAPIResponseModel).Reset()
}

// AlibabaWdkItemStoreskuUpdateAPIResponseModel is 五道口商品中心门店商品修改 成功返回结果
type AlibabaWdkItemStoreskuUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storesku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemStoreskuUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemStoreskuUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemStoreskuUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemStoreskuUpdateAPIResponse)
	},
}

// GetAlibabaWdkItemStoreskuUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemStoreskuUpdateAPIResponse
func GetAlibabaWdkItemStoreskuUpdateAPIResponse() *AlibabaWdkItemStoreskuUpdateAPIResponse {
	return poolAlibabaWdkItemStoreskuUpdateAPIResponse.Get().(*AlibabaWdkItemStoreskuUpdateAPIResponse)
}

// ReleaseAlibabaWdkItemStoreskuUpdateAPIResponse 将 AlibabaWdkItemStoreskuUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemStoreskuUpdateAPIResponse(v *AlibabaWdkItemStoreskuUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemStoreskuUpdateAPIResponse.Put(v)
}
