package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateUpdateAPIResponse 修改运费模板 API返回值
// taobao.delivery.template.update
//
// 修改运费模板
type TaobaoDeliveryTemplateUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoDeliveryTemplateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeliveryTemplateUpdateAPIResponseModel).Reset()
}

// TaobaoDeliveryTemplateUpdateAPIResponseModel is 修改运费模板 成功返回结果
type TaobaoDeliveryTemplateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 表示修改是否成功
	Complete bool `json:"complete,omitempty" xml:"complete,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Complete = false
}

var poolTaobaoDeliveryTemplateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeliveryTemplateUpdateAPIResponse)
	},
}

// GetTaobaoDeliveryTemplateUpdateAPIResponse 从 sync.Pool 获取 TaobaoDeliveryTemplateUpdateAPIResponse
func GetTaobaoDeliveryTemplateUpdateAPIResponse() *TaobaoDeliveryTemplateUpdateAPIResponse {
	return poolTaobaoDeliveryTemplateUpdateAPIResponse.Get().(*TaobaoDeliveryTemplateUpdateAPIResponse)
}

// ReleaseTaobaoDeliveryTemplateUpdateAPIResponse 将 TaobaoDeliveryTemplateUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeliveryTemplateUpdateAPIResponse(v *TaobaoDeliveryTemplateUpdateAPIResponse) {
	v.Reset()
	poolTaobaoDeliveryTemplateUpdateAPIResponse.Put(v)
}
