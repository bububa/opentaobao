package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuWarehouseskuQueryAPIResponse 仓商品查询接口(指定商品编码) API返回值
// alibaba.wdk.sku.warehousesku.query
//
// 提供指定仓商品编码查询
type AlibabaWdkSkuWarehouseskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuWarehouseskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuWarehouseskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuWarehouseskuQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuWarehouseskuQueryAPIResponseModel is 仓商品查询接口(指定商品编码) 成功返回结果
type AlibabaWdkSkuWarehouseskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_warehousesku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaWdkSkuWarehouseskuQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuWarehouseskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuWarehouseskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuWarehouseskuQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuWarehouseskuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuWarehouseskuQueryAPIResponse
func GetAlibabaWdkSkuWarehouseskuQueryAPIResponse() *AlibabaWdkSkuWarehouseskuQueryAPIResponse {
	return poolAlibabaWdkSkuWarehouseskuQueryAPIResponse.Get().(*AlibabaWdkSkuWarehouseskuQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuWarehouseskuQueryAPIResponse 将 AlibabaWdkSkuWarehouseskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuWarehouseskuQueryAPIResponse(v *AlibabaWdkSkuWarehouseskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuWarehouseskuQueryAPIResponse.Put(v)
}
