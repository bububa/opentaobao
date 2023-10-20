package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateGetAPIResponse 获取用户指定运费模板信息 API返回值
// taobao.delivery.template.get
//
// 获取用户指定运费模板信息
type TaobaoDeliveryTemplateGetAPIResponse struct {
	model.CommonResponse
	TaobaoDeliveryTemplateGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeliveryTemplateGetAPIResponseModel).Reset()
}

// TaobaoDeliveryTemplateGetAPIResponseModel is 获取用户指定运费模板信息 成功返回结果
type TaobaoDeliveryTemplateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 运费模板列表
	DeliveryTemplates []DeliveryTemplate `json:"delivery_templates,omitempty" xml:"delivery_templates>delivery_template,omitempty"`
	// 获得到符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliveryTemplates = m.DeliveryTemplates[:0]
	m.TotalResults = 0
}

var poolTaobaoDeliveryTemplateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeliveryTemplateGetAPIResponse)
	},
}

// GetTaobaoDeliveryTemplateGetAPIResponse 从 sync.Pool 获取 TaobaoDeliveryTemplateGetAPIResponse
func GetTaobaoDeliveryTemplateGetAPIResponse() *TaobaoDeliveryTemplateGetAPIResponse {
	return poolTaobaoDeliveryTemplateGetAPIResponse.Get().(*TaobaoDeliveryTemplateGetAPIResponse)
}

// ReleaseTaobaoDeliveryTemplateGetAPIResponse 将 TaobaoDeliveryTemplateGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeliveryTemplateGetAPIResponse(v *TaobaoDeliveryTemplateGetAPIResponse) {
	v.Reset()
	poolTaobaoDeliveryTemplateGetAPIResponse.Put(v)
}
