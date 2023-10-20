package openmall

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallRefundCloseAPIRequest) Reset() {
	r._distributor = ""
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundCloseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallRefundCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道
func (r *TaobaoOpenmallRefundCloseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundCloseAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundId is RefundId Setter
// 退款ID
func (r *TaobaoOpenmallRefundCloseAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundCloseAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoOpenmallRefundCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallRefundCloseRequest()
	},
}

// GetTaobaoOpenmallRefundCloseRequest 从 sync.Pool 获取 TaobaoOpenmallRefundCloseAPIRequest
func GetTaobaoOpenmallRefundCloseAPIRequest() *TaobaoOpenmallRefundCloseAPIRequest {
	return poolTaobaoOpenmallRefundCloseAPIRequest.Get().(*TaobaoOpenmallRefundCloseAPIRequest)
}

// ReleaseTaobaoOpenmallRefundCloseAPIRequest 将 TaobaoOpenmallRefundCloseAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallRefundCloseAPIRequest(v *TaobaoOpenmallRefundCloseAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallRefundCloseAPIRequest.Put(v)
}
