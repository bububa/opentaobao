package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionMaterialCreateAPIResponse 小程序投放--新建投放素材 API返回值
// taobao.miniapp.distribution.material.create
//
// 为可投放的小程序，增加入口的素材信息，比如图片、引导文案等等。
type TaobaoMiniappDistributionMaterialCreateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionMaterialCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionMaterialCreateAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionMaterialCreateAPIResponseModel is 小程序投放--新建投放素材 成功返回结果
type TaobaoMiniappDistributionMaterialCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_material_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MaterialErrorCode string `json:"material_error_code,omitempty" xml:"material_error_code,omitempty"`
	// 错误信息
	MaterialErrorMessage string `json:"material_error_message,omitempty" xml:"material_error_message,omitempty"`
	// 创建成功的素材id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	MaterialSuccess bool `json:"material_success,omitempty" xml:"material_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionMaterialCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MaterialErrorCode = ""
	m.MaterialErrorMessage = ""
	m.Model = 0
	m.MaterialSuccess = false
}

var poolTaobaoMiniappDistributionMaterialCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionMaterialCreateAPIResponse)
	},
}

// GetTaobaoMiniappDistributionMaterialCreateAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionMaterialCreateAPIResponse
func GetTaobaoMiniappDistributionMaterialCreateAPIResponse() *TaobaoMiniappDistributionMaterialCreateAPIResponse {
	return poolTaobaoMiniappDistributionMaterialCreateAPIResponse.Get().(*TaobaoMiniappDistributionMaterialCreateAPIResponse)
}

// ReleaseTaobaoMiniappDistributionMaterialCreateAPIResponse 将 TaobaoMiniappDistributionMaterialCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionMaterialCreateAPIResponse(v *TaobaoMiniappDistributionMaterialCreateAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionMaterialCreateAPIResponse.Put(v)
}
