package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse 小部件实例化版本更新 API返回值
// taobao.miniapp.widget.template.instance.update
//
// 小部件版本更新
type TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponseModel).Reset()
}

// TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponseModel is 小部件实例化版本更新 成功返回结果
type TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_widget_template_instance_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	Result *TaobaoMiniappWidgetTemplateInstanceUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse)
	},
}

// GetTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse 从 sync.Pool 获取 TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse
func GetTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse() *TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse {
	return poolTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse.Get().(*TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse)
}

// ReleaseTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse 将 TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse(v *TaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoMiniappWidgetTemplateInstanceUpdateAPIResponse.Put(v)
}
