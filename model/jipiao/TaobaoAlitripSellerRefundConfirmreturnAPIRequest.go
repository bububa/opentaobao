package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripSellerRefundConfirmreturnAPIRequest
【机票代理商】确认退票 API请求
taobao.alitrip.seller.refund.confirmreturn

确认退票 */
type TaobaoAlitripSellerRefundConfirmreturnAPIRequest struct {
	model.Params
	// 退票申请单
	_applyId int64
}

// NewTaobaoAlitripSellerRefundConfirmreturnRequest 初始化TaobaoAlitripSellerRefundConfirmreturnAPIRequest对象
func NewTaobaoAlitripSellerRefundConfirmreturnRequest() *TaobaoAlitripSellerRefundConfirmreturnAPIRequest {
	return &TaobaoAlitripSellerRefundConfirmreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.seller.refund.confirmreturn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ApplyId Setter
// 退票申请单
func (r *TaobaoAlitripSellerRefundConfirmreturnAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// Get ApplyId Getter
func (r TaobaoAlitripSellerRefundConfirmreturnAPIRequest) GetApplyId() int64 {
	return r._applyId
}
