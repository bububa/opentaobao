package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundgetAPIRequest 获取OpenMall退款单详情 API请求
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
type TaobaoopenmallrefundgetAPIRequest struct {
	model.Params
	// 渠道商身份
	_distributor string
	// 退款单ID
	_refundId int64
}

// NewTaobaoopenmallrefundgetRequest 初始化TaobaoopenmallrefundgetAPIRequest对象
func NewTaobaoopenmallrefundgetRequest() *TaobaoopenmallrefundgetAPIRequest {
	return &TaobaoopenmallrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道商身份
func (r *TaobaoopenmallrefundgetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundgetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoopenmallrefundgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
