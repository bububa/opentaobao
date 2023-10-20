package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialGetAPIResponse 小程序投放---读取投放入口素材信息 API返回值
// taobao.miniapp.distribution.material.get
//
// 读取已录入的投放入口素材信息。
type TaobaoMiniappDistributionMaterialGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionMaterialGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionMaterialGetAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionMaterialGetAPIResponseModel is 小程序投放---读取投放入口素材信息 成功返回结果
type TaobaoMiniappDistributionMaterialGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_material_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 素材信息
	Model []TaobaoMiniappDistributionMaterialGetModel `json:"model,omitempty" xml:"model>taobao_miniapp_distribution_material_get_model,omitempty"`
	// 错误码
	MaterialErrorCode string `json:"material_error_code,omitempty" xml:"material_error_code,omitempty"`
	// 错误信息
	MaterialErrorMessage string `json:"material_error_message,omitempty" xml:"material_error_message,omitempty"`
	// 调用是否成功
	MaterialSuccess bool `json:"material_success,omitempty" xml:"material_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = m.Model[:0]
	m.MaterialErrorCode = ""
	m.MaterialErrorMessage = ""
	m.MaterialSuccess = false
}

var poolTaobaoMiniappDistributionMaterialGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionMaterialGetAPIResponse)
	},
}

// GetTaobaoMiniappDistributionMaterialGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialGetAPIResponse
func GetTaobaoMiniappDistributionMaterialGetAPIResponse() *TaobaoMiniappDistributionMaterialGetAPIResponse {
	return poolTaobaoMiniappDistributionMaterialGetAPIResponse.Get().(*TaobaoMiniappDistributionMaterialGetAPIResponse)
}

// ReleaseTaobaoMiniappDistributionMaterialGetAPIResponse 将 TaobaoMiniappDistributionMaterialGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialGetAPIResponse(v *TaobaoMiniappDistributionMaterialGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialGetAPIResponse.Put(v)
}
