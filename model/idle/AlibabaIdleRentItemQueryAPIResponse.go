package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemQueryAPIResponse 查询租赁商品信息 API返回值
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
type AlibabaIdleRentItemQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentItemQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentItemQueryAPIResponseModel).Reset()
}

// AlibabaIdleRentItemQueryAPIResponseModel is 查询租赁商品信息 成功返回结果
type AlibabaIdleRentItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleRentItemQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentItemQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemQueryAPIResponse)
	},
}

// GetAlibabaIdleRentItemQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleRentItemQueryAPIResponse
func GetAlibabaIdleRentItemQueryAPIResponse() *AlibabaIdleRentItemQueryAPIResponse {
	return poolAlibabaIdleRentItemQueryAPIResponse.Get().(*AlibabaIdleRentItemQueryAPIResponse)
}

// ReleaseAlibabaIdleRentItemQueryAPIResponse 将 AlibabaIdleRentItemQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentItemQueryAPIResponse(v *AlibabaIdleRentItemQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentItemQueryAPIResponse.Put(v)
}
