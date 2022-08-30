package alidoc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchquery 商家维度批量查询订单处方详情
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
func AlibabaAlihealthSellerRxPrescriptionDetailBatchquery(clt *core.SDKClient, req *alidoc.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest, session string) (*alidoc.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse, error) {
	var resp alidoc.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
