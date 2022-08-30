package miniappopen

import (
	"encoding/xml"

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
