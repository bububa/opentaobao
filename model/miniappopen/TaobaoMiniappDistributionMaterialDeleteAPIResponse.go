package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialDeleteAPIResponse 小程序投放 --- 删除投放素材 API返回值
// taobao.miniapp.distribution.material.delete
//
// 删除已录入的投放入口素材信息。
type TaobaoMiniappDistributionMaterialDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionMaterialDeleteAPIResponseModel
}

// TaobaoMiniappDistributionMaterialDeleteAPIResponseModel is 小程序投放 --- 删除投放素材 成功返回结果
type TaobaoMiniappDistributionMaterialDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_material_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MaterialErrorCode string `json:"material_error_code,omitempty" xml:"material_error_code,omitempty"`
	// 错误信息
	MaterialErrorMessage string `json:"material_error_message,omitempty" xml:"material_error_message,omitempty"`
	// 调用是否成功
	MaterialSuccess bool `json:"material_success,omitempty" xml:"material_success,omitempty"`
	// 操作结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
