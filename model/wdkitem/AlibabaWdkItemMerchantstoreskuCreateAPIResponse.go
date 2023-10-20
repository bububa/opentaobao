package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemMerchantstoreskuCreateAPIResponse 门店商品信息新建 API返回值
// alibaba.wdk.item.merchantstoresku.create
//
// 门店商品信息新建
type AlibabaWdkItemMerchantstoreskuCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantstoreskuCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel).Reset()
}

// AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel is 门店商品信息新建 成功返回结果
type AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_merchantstoresku_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemMerchantstoreskuCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemMerchantstoreskuCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemMerchantstoreskuCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemMerchantstoreskuCreateAPIResponse)
	},
}

// GetAlibabaWdkItemMerchantstoreskuCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkItemMerchantstoreskuCreateAPIResponse
func GetAlibabaWdkItemMerchantstoreskuCreateAPIResponse() *AlibabaWdkItemMerchantstoreskuCreateAPIResponse {
	return poolAlibabaWdkItemMerchantstoreskuCreateAPIResponse.Get().(*AlibabaWdkItemMerchantstoreskuCreateAPIResponse)
}

// ReleaseAlibabaWdkItemMerchantstoreskuCreateAPIResponse 将 AlibabaWdkItemMerchantstoreskuCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemMerchantstoreskuCreateAPIResponse(v *AlibabaWdkItemMerchantstoreskuCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemMerchantstoreskuCreateAPIResponse.Put(v)
}
