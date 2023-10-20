package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockinsurancerefundgetAPIRequest 共享库存投保业务售后逆向订单数据获取 API请求
// alibaba.wdkorder.sharestock.insurance.refundget
//
// 共享库存投保业务售后逆向订单数据获取
type AlibabawdkordersharestockinsurancerefundgetAPIRequest struct {
	model.Params
	// 淘宝子订单ID
	_tbSubOrderId string
	// 退货单ID
	_refundId string
}

// NewAlibabawdkordersharestockinsurancerefundgetRequest 初始化AlibabawdkordersharestockinsurancerefundgetAPIRequest对象
func NewAlibabawdkordersharestockinsurancerefundgetRequest() *AlibabawdkordersharestockinsurancerefundgetAPIRequest {
	return &AlibabawdkordersharestockinsurancerefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersharestockinsurancerefundgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.insurance.refundget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersharestockinsurancerefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersharestockinsurancerefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTbSubOrderId is TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabawdkordersharestockinsurancerefundgetAPIRequest) SetTbSubOrderId(_tbSubOrderId string) error {
	r._tbSubOrderId = _tbSubOrderId
	r.Set("tb_sub_order_id", _tbSubOrderId)
	return nil
}

// GetTbSubOrderId TbSubOrderId Getter
func (r AlibabawdkordersharestockinsurancerefundgetAPIRequest) GetTbSubOrderId() string {
	return r._tbSubOrderId
}

// SetRefundId is RefundId Setter
// 退货单ID
func (r *AlibabawdkordersharestockinsurancerefundgetAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabawdkordersharestockinsurancerefundgetAPIRequest) GetRefundId() string {
	return r._refundId
}
