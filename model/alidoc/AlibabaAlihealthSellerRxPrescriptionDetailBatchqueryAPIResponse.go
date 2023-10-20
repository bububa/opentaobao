package alidoc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse 商家维度批量查询订单处方详情 API返回值
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
type AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponseModel).Reset()
}

// AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponseModel is 商家维度批量查询订单处方详情 成功返回结果
type AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_seller_rx_prescription_detail_batchquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse)
	},
}

// GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse 从 sync.Pool 获取 AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse
func GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse() *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse {
	return poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse.Get().(*AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse)
}

// ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse 将 AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse(v *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse.Put(v)
}
