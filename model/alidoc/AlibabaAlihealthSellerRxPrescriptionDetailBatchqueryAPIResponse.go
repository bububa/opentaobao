package alidoc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIResponse 商家维度批量查询订单处方详情 API返回值
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
type AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIResponseModel
}

// AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIResponseModel is 商家维度批量查询订单处方详情 成功返回结果
type AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_seller_rx_prescription_detail_batchquery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihealthsellerrxprescriptiondetailbatchqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
