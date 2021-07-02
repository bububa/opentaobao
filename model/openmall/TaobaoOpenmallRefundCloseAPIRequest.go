package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundCloseAPIRequest 关闭OpenMall退款单 API请求
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
type TaobaoOpenmallRefundCloseAPIRequest struct {
	model.Params
	// 渠道
	_distributor string
	// 退款ID
	_refundId int64
}

// NewTaobaoOpenmallRefundCloseRequest 初始化TaobaoOpenmallRefundCloseAPIRequest对象
func NewTaobaoOpenmallRefundCloseRequest() *TaobaoOpenmallRefundCloseAPIRequest {
	return &TaobaoOpenmallRefundCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundCloseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Distributor Setter
// 渠道
func (r *TaobaoOpenmallRefundCloseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// Get Distributor Getter
func (r TaobaoOpenmallRefundCloseAPIRequest) GetDistributor() string {
	return r._distributor
}

// Set is RefundId Setter
// 退款ID
func (r *TaobaoOpenmallRefundCloseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r TaobaoOpenmallRefundCloseAPIRequest) GetRefundId() int64 {
	return r._refundId
}
