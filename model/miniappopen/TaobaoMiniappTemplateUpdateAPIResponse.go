package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappTemplateUpdateAPIResponse （已废弃）更新实例化应用 API返回值
// taobao.miniapp.template.update
//
// 商家应用c端模板实例化小程序更新
type TaobaoMiniappTemplateUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappTemplateUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappTemplateUpdateAPIResponseModel).Reset()
}

// TaobaoMiniappTemplateUpdateAPIResponseModel is （已废弃）更新实例化应用 成功返回结果
type TaobaoMiniappTemplateUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_template_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoMiniappTemplateUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappTemplateUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniappTemplateUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappTemplateUpdateAPIResponse)
	},
}

// GetTaobaoMiniappTemplateUpdateAPIResponse 从 sync.Pool 获取 TaobaoMiniappTemplateUpdateAPIResponse
func GetTaobaoMiniappTemplateUpdateAPIResponse() *TaobaoMiniappTemplateUpdateAPIResponse {
	return poolTaobaoMiniappTemplateUpdateAPIResponse.Get().(*TaobaoMiniappTemplateUpdateAPIResponse)
}

// ReleaseTaobaoMiniappTemplateUpdateAPIResponse 将 TaobaoMiniappTemplateUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappTemplateUpdateAPIResponse(v *TaobaoMiniappTemplateUpdateAPIResponse) {
	v.Reset()
	poolTaobaoMiniappTemplateUpdateAPIResponse.Put(v)
}
