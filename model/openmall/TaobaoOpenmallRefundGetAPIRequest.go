package openmall

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallRefundGetAPIRequest) Reset() {
	r._distributor = ""
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 渠道商身份
func (r *TaobaoOpenmallRefundGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallRefundGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *TaobaoOpenmallRefundGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoOpenmallRefundGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoOpenmallRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallRefundGetRequest()
	},
}

// GetTaobaoOpenmallRefundGetRequest 从 sync.Pool 获取 TaobaoOpenmallRefundGetAPIRequest
func GetTaobaoOpenmallRefundGetAPIRequest() *TaobaoOpenmallRefundGetAPIRequest {
	return poolTaobaoOpenmallRefundGetAPIRequest.Get().(*TaobaoOpenmallRefundGetAPIRequest)
}

// ReleaseTaobaoOpenmallRefundGetAPIRequest 将 TaobaoOpenmallRefundGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallRefundGetAPIRequest(v *TaobaoOpenmallRefundGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallRefundGetAPIRequest.Put(v)
}
