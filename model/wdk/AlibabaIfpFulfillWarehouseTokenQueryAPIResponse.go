package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIfpFulfillWarehouseTokenQueryAPIResponse 同城令牌打印接口 API返回值
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIfpFulfillWarehouseTokenQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel).Reset()
}

// AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel is 同城令牌打印接口 成功返回结果
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ifp_fulfill_warehouse_token_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	WorkResult *WorkResult `json:"work_result,omitempty" xml:"work_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIfpFulfillWarehouseTokenQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WorkResult = nil
}

var poolAlibabaIfpFulfillWarehouseTokenQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIfpFulfillWarehouseTokenQueryAPIResponse)
	},
}

// GetAlibabaIfpFulfillWarehouseTokenQueryAPIResponse 从 sync.Pool 获取 AlibabaIfpFulfillWarehouseTokenQueryAPIResponse
func GetAlibabaIfpFulfillWarehouseTokenQueryAPIResponse() *AlibabaIfpFulfillWarehouseTokenQueryAPIResponse {
	return poolAlibabaIfpFulfillWarehouseTokenQueryAPIResponse.Get().(*AlibabaIfpFulfillWarehouseTokenQueryAPIResponse)
}

// ReleaseAlibabaIfpFulfillWarehouseTokenQueryAPIResponse 将 AlibabaIfpFulfillWarehouseTokenQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIfpFulfillWarehouseTokenQueryAPIResponse(v *AlibabaIfpFulfillWarehouseTokenQueryAPIResponse) {
	v.Reset()
	poolAlibabaIfpFulfillWarehouseTokenQueryAPIResponse.Put(v)
}
