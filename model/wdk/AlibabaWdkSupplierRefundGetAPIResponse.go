package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSupplierRefundGetAPIResponse 五道口按订单号批量查询供应商退款单 API返回值
// alibaba.wdk.supplier.refund.get
//
// 五道口按订单号批量查询供应商退款单
type AlibabaWdkSupplierRefundGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierRefundGetAPIResponseModel
}

// AlibabaWdkSupplierRefundGetAPIResponseModel is 五道口按订单号批量查询供应商退款单 成功返回结果
type AlibabaWdkSupplierRefundGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_refund_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}
