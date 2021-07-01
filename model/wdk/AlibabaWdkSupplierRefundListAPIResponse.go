package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSupplierRefundListAPIResponse
五道口按供应商拉取退款单 API返回值
alibaba.wdk.supplier.refund.list

五道口按供应商拉取退款单 */
type AlibabaWdkSupplierRefundListAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSupplierRefundListAPIResponseModel
}

// AlibabaWdkSupplierRefundListAPIResponseModel is 五道口按供应商拉取退款单 成功返回结果
type AlibabaWdkSupplierRefundListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_supplier_refund_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *OrderSyncRefundListResult `json:"result,omitempty" xml:"result,omitempty"`
}
