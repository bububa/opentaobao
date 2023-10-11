package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsErpDeliveryCutAPIResponse ERP发起配拦截 API返回值
// taobao.logistics.erp.delivery.cut
//
// ERP发起配拦截
type TaobaoLogisticsErpDeliveryCutAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsErpDeliveryCutAPIResponseModel
}

// TaobaoLogisticsErpDeliveryCutAPIResponseModel is ERP发起配拦截 成功返回结果
type TaobaoLogisticsErpDeliveryCutAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_erp_delivery_cut_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	BizErrorMessage string `json:"biz_error_message,omitempty" xml:"biz_error_message,omitempty"`
	// 错误信息
	BizErrorCode string `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
	// 是否需要重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
