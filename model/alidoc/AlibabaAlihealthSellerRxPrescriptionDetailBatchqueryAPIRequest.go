package alidoc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest 商家维度批量查询订单处方详情 API请求
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
type AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest struct {
	model.Params
	// 订单号，请用逗号隔开
	_bizOrderIds string
}

// NewAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryRequest 初始化AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest对象
func NewAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryRequest() *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest {
	return &AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) Reset() {
	r._bizOrderIds = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.seller.rx.prescription.detail.batchquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderIds is BizOrderIds Setter
// 订单号，请用逗号隔开
func (r *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) SetBizOrderIds(_bizOrderIds string) error {
	r._bizOrderIds = _bizOrderIds
	r.Set("biz_order_ids", _bizOrderIds)
	return nil
}

// GetBizOrderIds BizOrderIds Getter
func (r AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) GetBizOrderIds() string {
	return r._bizOrderIds
}

var poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryRequest()
	},
}

// GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryRequest 从 sync.Pool 获取 AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest
func GetAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest() *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest {
	return poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest.Get().(*AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest)
}

// ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest 将 AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest(v *AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest.Put(v)
}
