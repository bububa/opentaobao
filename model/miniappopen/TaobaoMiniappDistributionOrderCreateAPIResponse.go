package miniappopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderCreateAPIResponse 创建小程序投放计划 API返回值
// taobao.miniapp.distribution.order.create
//
// 帮助商家，创建小程序的投放计划。该api，仅针对特定场景开放，目前仅支持客服场景，具体支持的场景列表请联系技术支持或业务对接人确认。
type TaobaoMiniappDistributionOrderCreateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappDistributionOrderCreateAPIResponseModel).Reset()
}

// TaobaoMiniappDistributionOrderCreateAPIResponseModel is 创建小程序投放计划 成功返回结果
type TaobaoMiniappDistributionOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	OrderErrorInfo string `json:"order_error_info,omitempty" xml:"order_error_info,omitempty"`
	// 错误码
	OrderErrorCode int64 `json:"order_error_code,omitempty" xml:"order_error_code,omitempty"`
	// 创建成功的投放计划id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	OrderSuccess bool `json:"order_success,omitempty" xml:"order_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappDistributionOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderErrorInfo = ""
	m.OrderErrorCode = 0
	m.Model = 0
	m.OrderSuccess = false
}

var poolTaobaoMiniappDistributionOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappDistributionOrderCreateAPIResponse)
	},
}

// GetTaobaoMiniappDistributionOrderCreateAPIResponse 从 sync.Pool 获取 TaobaoMiniappDistributionOrderCreateAPIResponse
func GetTaobaoMiniappDistributionOrderCreateAPIResponse() *TaobaoMiniappDistributionOrderCreateAPIResponse {
	return poolTaobaoMiniappDistributionOrderCreateAPIResponse.Get().(*TaobaoMiniappDistributionOrderCreateAPIResponse)
}

// ReleaseTaobaoMiniappDistributionOrderCreateAPIResponse 将 TaobaoMiniappDistributionOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappDistributionOrderCreateAPIResponse(v *TaobaoMiniappDistributionOrderCreateAPIResponse) {
	v.Reset()
	poolTaobaoMiniappDistributionOrderCreateAPIResponse.Put(v)
}
