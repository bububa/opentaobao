package miniappopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappDistributionOrderPrecreateAPIResponse 代商家预创建投放计划 API返回值
// taobao.miniapp.distribution.order.precreate
//
// 帮助商家，预创建小程序的投放计划，预创建的投放计划，在商家确认以后，则会生效可用。
type TaobaoMiniappDistributionOrderPrecreateAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappDistributionOrderPrecreateAPIResponseModel
}

// TaobaoMiniappDistributionOrderPrecreateAPIResponseModel is 代商家预创建投放计划 成功返回结果
type TaobaoMiniappDistributionOrderPrecreateAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_distribution_order_precreate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	OrderErrorInfo string `json:"order_error_info,omitempty" xml:"order_error_info,omitempty"`
	// 错误码
	OrderErrorCode int64 `json:"order_error_code,omitempty" xml:"order_error_code,omitempty"`
	// 预创建成功的投放计划id
	Model int64 `json:"model,omitempty" xml:"model,omitempty"`
	// 调用是否成功
	OrderSuccess bool `json:"order_success,omitempty" xml:"order_success,omitempty"`
}
