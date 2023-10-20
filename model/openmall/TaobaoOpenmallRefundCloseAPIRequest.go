package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundcloseAPIRequest 关闭OpenMall退款单 API请求
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
type TaobaoopenmallrefundcloseAPIRequest struct {
	model.Params
	// 渠道
	_distributor string
	// 退款ID
	_refundId int64
}

// NewTaobaoopenmallrefundcloseRequest 初始化TaobaoopenmallrefundcloseAPIRequest对象
func NewTaobaoopenmallrefundcloseRequest() *TaobaoopenmallrefundcloseAPIRequest {
	return &TaobaoopenmallrefundcloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmallrefundcloseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmallrefundcloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmallrefundcloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道
func (r *TaobaoopenmallrefundcloseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmallrefundcloseAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundId is RefundId Setter
// 退款ID
func (r *TaobaoopenmallrefundcloseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoopenmallrefundcloseAPIRequest) GetRefundId() int64 {
	return r._refundId
}
