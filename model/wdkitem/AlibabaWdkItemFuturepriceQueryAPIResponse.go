package wdkitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemFuturepriceQueryAPIResponse 单个商品未来价查询接口 API返回值
// alibaba.wdk.item.futureprice.query
//
// 查询单个商品未来价，融合了未来基础售价+未来促销价
type AlibabaWdkItemFuturepriceQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemFuturepriceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemFuturepriceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemFuturepriceQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemFuturepriceQueryAPIResponseModel is 单个商品未来价查询接口 成功返回结果
type AlibabaWdkItemFuturepriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_futureprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaWdkItemFuturepriceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemFuturepriceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemFuturepriceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemFuturepriceQueryAPIResponse)
	},
}

// GetAlibabaWdkItemFuturepriceQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemFuturepriceQueryAPIResponse
func GetAlibabaWdkItemFuturepriceQueryAPIResponse() *AlibabaWdkItemFuturepriceQueryAPIResponse {
	return poolAlibabaWdkItemFuturepriceQueryAPIResponse.Get().(*AlibabaWdkItemFuturepriceQueryAPIResponse)
}

// ReleaseAlibabaWdkItemFuturepriceQueryAPIResponse 将 AlibabaWdkItemFuturepriceQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemFuturepriceQueryAPIResponse(v *AlibabaWdkItemFuturepriceQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemFuturepriceQueryAPIResponse.Put(v)
}
