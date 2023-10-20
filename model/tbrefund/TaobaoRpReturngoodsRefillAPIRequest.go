package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRpReturngoodsRefillAPIRequest 卖家回填物流信息 API请求
// taobao.rp.returngoods.refill
//
// 卖家收到货物回填物流信息，如果买家已经回填物流信息，则接口报错，目前仅支持天猫订单。
type TaobaoRpReturngoodsRefillAPIRequest struct {
	model.Params
	// 退款阶段，可选值：售中：onsale，售后：aftersale
	_refundPhase string
	// 物流公司运单号
	_logisticsWaybillNo string
	// 物流公司编号
	_logisticsCompanyCode string
	// 退款单编号
	_refundId int64
}

// NewTaobaoRpReturngoodsRefillRequest 初始化TaobaoRpReturngoodsRefillAPIRequest对象
func NewTaobaoRpReturngoodsRefillRequest() *TaobaoRpReturngoodsRefillAPIRequest {
	return &TaobaoRpReturngoodsRefillAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRpReturngoodsRefillAPIRequest) Reset() {
	r._refundPhase = ""
	r._logisticsWaybillNo = ""
	r._logisticsCompanyCode = ""
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsRefillAPIRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsRefillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRpReturngoodsRefillAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoRpReturngoodsRefillAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRpReturngoodsRefillRequest()
	},
}

// GetTaobaoRpReturngoodsRefillRequest 从 sync.Pool 获取 TaobaoRpReturngoodsRefillAPIRequest
func GetTaobaoRpReturngoodsRefillAPIRequest() *TaobaoRpReturngoodsRefillAPIRequest {
	return poolTaobaoRpReturngoodsRefillAPIRequest.Get().(*TaobaoRpReturngoodsRefillAPIRequest)
}

// ReleaseTaobaoRpReturngoodsRefillAPIRequest 将 TaobaoRpReturngoodsRefillAPIRequest 放入 sync.Pool
func ReleaseTaobaoRpReturngoodsRefillAPIRequest(v *TaobaoRpReturngoodsRefillAPIRequest) {
	v.Reset()
	poolTaobaoRpReturngoodsRefillAPIRequest.Put(v)
}
