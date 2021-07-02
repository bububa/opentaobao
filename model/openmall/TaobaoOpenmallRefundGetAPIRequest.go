package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallRefundGetAPIRequest 获取OpenMall退款单详情 API请求
// taobao.openmall.refund.get
//
// 获取OpenMall退款单详情
type TaobaoOpenmallRefundGetAPIRequest struct {
	model.Params
	// 渠道商身份
	_distributor string
	// 退款单ID
	_refundId int64
}

// NewTaobaoOpenmallRefundGetRequest 初始化TaobaoOpenmallRefundGetAPIRequest对象
func NewTaobaoOpenmallRefundGetRequest() *TaobaoOpenmallRefundGetAPIRequest {
	return &TaobaoOpenmallRefundGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Distributor Setter
// 渠道商身份
func (r *TaobaoOpenmallRefundGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// Get Distributor Getter
func (r TaobaoOpenmallRefundGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// Set is RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r TaobaoOpenmallRefundGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
