package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsRefillAPIRequest 卖家回填物流信息 API请求
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillAPIRequest struct {
	model.Params
	// 退款单编号
	_refundId int64
	// 退款阶段，可选值：售中：onsale，售后：aftersale
	_refundPhase string
	// 物流公司运单号
	_logisticsWaybillNo string
	// 物流公司编号
	_logisticsCompanyCode string
}

// NewTaobaoRpReturngoodsRefillRequest 初始化TaobaoRpReturngoodsRefillAPIRequest对象
func NewTaobaoRpReturngoodsRefillRequest() *TaobaoRpReturngoodsRefillAPIRequest {
	return &TaobaoRpReturngoodsRefillAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefillAPIRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefillAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefundId is RefundId Setter
// 退款单编号
func (r *TaobaoRpReturngoodsRefillAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRpReturngoodsRefillAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetRefundPhase is RefundPhase Setter
// 退款阶段，可选值：售中：onsale，售后：aftersale
func (r *TaobaoRpReturngoodsRefillAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaoRpReturngoodsRefillAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetLogisticsWaybillNo is LogisticsWaybillNo Setter
// 物流公司运单号
func (r *TaobaoRpReturngoodsRefillAPIRequest) SetLogisticsWaybillNo(_logisticsWaybillNo string) error {
	r._logisticsWaybillNo = _logisticsWaybillNo
	r.Set("logistics_waybill_no", _logisticsWaybillNo)
	return nil
}

// GetLogisticsWaybillNo LogisticsWaybillNo Getter
func (r TaobaoRpReturngoodsRefillAPIRequest) GetLogisticsWaybillNo() string {
	return r._logisticsWaybillNo
}

// SetLogisticsCompanyCode is LogisticsCompanyCode Setter
// 物流公司编号
func (r *TaobaoRpReturngoodsRefillAPIRequest) SetLogisticsCompanyCode(_logisticsCompanyCode string) error {
	r._logisticsCompanyCode = _logisticsCompanyCode
	r.Set("logistics_company_code", _logisticsCompanyCode)
	return nil
}

// GetLogisticsCompanyCode LogisticsCompanyCode Getter
func (r TaobaoRpReturngoodsRefillAPIRequest) GetLogisticsCompanyCode() string {
	return r._logisticsCompanyCode
}
