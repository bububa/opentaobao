package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniapppTemplateInstantiateAPIResponse （已废弃）构建实例化应用 API返回值
// taobao.miniappp.template.instantiate
//
// 实例化saas化的小程序
type TaobaoMiniapppTemplateInstantiateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniapppTemplateInstantiateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniapppTemplateInstantiateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniapppTemplateInstantiateAPIResponseModel).Reset()
}

// TaobaoMiniapppTemplateInstantiateAPIResponseModel is （已废弃）构建实例化应用 成功返回结果
type TaobaoMiniapppTemplateInstantiateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniappp_template_instantiate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoMiniapppTemplateInstantiateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniapppTemplateInstantiateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoMiniapppTemplateInstantiateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniapppTemplateInstantiateAPIResponse)
	},
}

// GetTaobaoMiniapppTemplateInstantiateAPIResponse 从 sync.Pool 获取 TaobaoMiniapppTemplateInstantiateAPIResponse
func GetTaobaoMiniapppTemplateInstantiateAPIResponse() *TaobaoMiniapppTemplateInstantiateAPIResponse {
	return poolTaobaoMiniapppTemplateInstantiateAPIResponse.Get().(*TaobaoMiniapppTemplateInstantiateAPIResponse)
}

// ReleaseTaobaoMiniapppTemplateInstantiateAPIResponse 将 TaobaoMiniapppTemplateInstantiateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniapppTemplateInstantiateAPIResponse(v *TaobaoMiniapppTemplateInstantiateAPIResponse) {
	v.Reset()
	poolTaobaoMiniapppTemplateInstantiateAPIResponse.Put(v)
}
