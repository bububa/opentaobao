package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalLogisticOrderCreateAPIResponse 创建物流订单 API返回值
// cainiao.global.logistic.order.create
//
// 创建物流订单
type CainiaoGlobalLogisticOrderCreateAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalLogisticOrderCreateAPIResponseModel
}

// CainiaoGlobalLogisticOrderCreateAPIResponseModel is 创建物流订单 成功返回结果
type CainiaoGlobalLogisticOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_logistic_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建是否成功
	IsSuccess string `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 请求结果
	Result *OpenTakingOrderResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 错误信息
	ErrorInfo *ErrorInfo `json:"error_info,omitempty" xml:"error_info,omitempty"`
	// 是否需要重试
	NeedRetry bool `json:"need_retry,omitempty" xml:"need_retry,omitempty"`
}
