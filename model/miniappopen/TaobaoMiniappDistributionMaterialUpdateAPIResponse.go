package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialUpdateAPIResponse 小程序投放 --- 更新投放素材 API返回值
// taobao.miniapp.distribution.material.update
//
// 修改已录入的投放素材信息。
type TaobaoMiniappDistributionMaterialUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionMaterialUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionMaterialUpdateAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionMaterialUpdateAPIResponseModel is 小程序投放 --- 更新投放素材 成功返回结果
type TaobaoMiniappDistributionMaterialUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_material_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MaterialErrorCode string `json:"material_error_code,omitempty" xml:"material_error_code,omitempty"`
	// 错误信息
	MaterialErrorMessage string `json:"material_error_message,omitempty" xml:"material_error_message,omitempty"`
	// 调用是否成功
	MaterialSuccess bool `json:"material_success,omitempty" xml:"material_success,omitempty"`
	// 是否修改成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MaterialErrorCode = ""
	m.MaterialErrorMessage = ""
	m.MaterialSuccess = false
	m.Model = false
}

var poolTaobaoMiniappDistributionMaterialUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionMaterialUpdateAPIResponse)
	},
}

// GetTaobaoMiniappDistributionMaterialUpdateAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialUpdateAPIResponse
func GetTaobaoMiniappDistributionMaterialUpdateAPIResponse() *TaobaoMiniappDistributionMaterialUpdateAPIResponse {
	return poolTaobaoMiniappDistributionMaterialUpdateAPIResponse.Get().(*TaobaoMiniappDistributionMaterialUpdateAPIResponse)
}

// ReleaseTaobaoMiniappDistributionMaterialUpdateAPIResponse 将 TaobaoMiniappDistributionMaterialUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialUpdateAPIResponse(v *TaobaoMiniappDistributionMaterialUpdateAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialUpdateAPIResponse.Put(v)
}
