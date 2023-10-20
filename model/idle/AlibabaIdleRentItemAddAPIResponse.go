package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemAddAPIResponse 租赁商品发布 API返回值
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
type AlibabaIdleRentItemAddAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentItemAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentItemAddAPIResponseModel).Reset()
}

// AlibabaIdleRentItemAddAPIResponseModel is 租赁商品发布 成功返回结果
type AlibabaIdleRentItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentItemAddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentItemAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemAddAPIResponse)
	},
}

// GetAlibabaIdleRentItemAddAPIResponse 从 sync.Pool 获取 AlibabaIdleRentItemAddAPIResponse
func GetAlibabaIdleRentItemAddAPIResponse() *AlibabaIdleRentItemAddAPIResponse {
	return poolAlibabaIdleRentItemAddAPIResponse.Get().(*AlibabaIdleRentItemAddAPIResponse)
}

// ReleaseAlibabaIdleRentItemAddAPIResponse 将 AlibabaIdleRentItemAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentItemAddAPIResponse(v *AlibabaIdleRentItemAddAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentItemAddAPIResponse.Put(v)
}
