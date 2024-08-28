package alidoc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alidoc"
)

// AlibabaAlihealthSellerRxPrescriptionDetailBatchquery 商家维度批量查询订单处方详情
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
func AlibabaAlihealthSellerRxPrescriptionDetailBatchquery(ctx context.Context, clt *core.SDKClient, req *alidoc.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIRequest, resp *alidoc.AlibabaAlihealthSellerRxPrescriptionDetailBatchqueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
