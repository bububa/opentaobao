package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentItemEditAPIResponse 租赁商品编辑 API返回值
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
type AlibabaIdleRentItemEditAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRentItemEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleRentItemEditAPIResponseModel).Reset()
}

// AlibabaIdleRentItemEditAPIResponseModel is 租赁商品编辑 成功返回结果
type AlibabaIdleRentItemEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_item_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaIdleRentItemEditTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleRentItemEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleRentItemEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentItemEditAPIResponse)
	},
}

// GetAlibabaIdleRentItemEditAPIResponse 从 sync.Pool 获取 AlibabaIdleRentItemEditAPIResponse
func GetAlibabaIdleRentItemEditAPIResponse() *AlibabaIdleRentItemEditAPIResponse {
	return poolAlibabaIdleRentItemEditAPIResponse.Get().(*AlibabaIdleRentItemEditAPIResponse)
}

// ReleaseAlibabaIdleRentItemEditAPIResponse 将 AlibabaIdleRentItemEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleRentItemEditAPIResponse(v *AlibabaIdleRentItemEditAPIResponse) {
	v.Reset()
	poolAlibabaIdleRentItemEditAPIResponse.Put(v)
}
