package alidoc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest 商家维度批量查询订单处方详情 API请求
// alibaba.alihealth.seller.rx.prescription.detail.batchquery
//
// 商家可以通过seller_id、biz_order_ids，批量查询自家店铺订单对应的处方详情。
type AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest struct {
	model.Params
	// 订单号，请用逗号隔开
	_bizOrderIds string
}

// NewAlibabaalihealthsellerrxprescriptiondetailbatchqueryRequest 初始化AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest对象
func NewAlibabaalihealthsellerrxprescriptiondetailbatchqueryRequest() *AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest {
	return &AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.seller.rx.prescription.detail.batchquery"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizOrderIds is BizOrderIds Setter
// 订单号，请用逗号隔开
func (r *AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest) SetBizOrderIds(_bizOrderIds string) error {
	r._bizOrderIds = _bizOrderIds
	r.Set("biz_order_ids", _bizOrderIds)
	return nil
}

// GetBizOrderIds BizOrderIds Getter
func (r AlibabaalihealthsellerrxprescriptiondetailbatchqueryAPIRequest) GetBizOrderIds() string {
	return r._bizOrderIds
}
