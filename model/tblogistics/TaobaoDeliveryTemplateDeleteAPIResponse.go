package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeliveryTemplateDeleteAPIResponse 删除运费模板 API返回值
// taobao.delivery.template.delete
//
// 根据用户指定的模板ID删除指定的模板
type TaobaoDeliveryTemplateDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoDeliveryTemplateDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeliveryTemplateDeleteAPIResponseModel).Reset()
}

// TaobaoDeliveryTemplateDeleteAPIResponseModel is 删除运费模板 成功返回结果
type TaobaoDeliveryTemplateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"delivery_template_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 表示删除成功还是失败
	Complete bool `json:"complete,omitempty" xml:"complete,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeliveryTemplateDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Complete = false
}

var poolTaobaoDeliveryTemplateDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeliveryTemplateDeleteAPIResponse)
	},
}

// GetTaobaoDeliveryTemplateDeleteAPIResponse 从 sync.Pool 获取 TaobaoDeliveryTemplateDeleteAPIResponse
func GetTaobaoDeliveryTemplateDeleteAPIResponse() *TaobaoDeliveryTemplateDeleteAPIResponse {
	return poolTaobaoDeliveryTemplateDeleteAPIResponse.Get().(*TaobaoDeliveryTemplateDeleteAPIResponse)
}

// ReleaseTaobaoDeliveryTemplateDeleteAPIResponse 将 TaobaoDeliveryTemplateDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeliveryTemplateDeleteAPIResponse(v *TaobaoDeliveryTemplateDeleteAPIResponse) {
	v.Reset()
	poolTaobaoDeliveryTemplateDeleteAPIResponse.Put(v)
}
