package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest 回收商同步前置补贴红包的代扣成功事件 API请求
// taobao.recycle.ofnpreredpacket.tpdeductsuccess
//
// 回收商-&gt;天猫后端，同步前置补贴红包的代扣成功事件
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest struct {
	model.Params
	// 变化的金额
	_deductAmount int64
	// 旧机单id
	_oldOrderId int64
}

// NewTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest 初始化TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest对象
func NewTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest() *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest {
	return &TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) Reset() {
	r._deductAmount = 0
	r._oldOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetApiMethodName() string {
	return "taobao.recycle.ofnpreredpacket.tpdeductsuccess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeductAmount is DeductAmount Setter
// 变化的金额
func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) SetDeductAmount(_deductAmount int64) error {
	r._deductAmount = _deductAmount
	r.Set("deduct_amount", _deductAmount)
	return nil
}

// GetDeductAmount DeductAmount Getter
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetDeductAmount() int64 {
	return r._deductAmount
}

// SetOldOrderId is OldOrderId Setter
// 旧机单id
func (r *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) SetOldOrderId(_oldOrderId int64) error {
	r._oldOrderId = _oldOrderId
	r.Set("old_order_id", _oldOrderId)
	return nil
}

// GetOldOrderId OldOrderId Getter
func (r TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) GetOldOrderId() int64 {
	return r._oldOrderId
}

var poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest()
	},
}

// GetTaobaoRecycleOfnpreredpacketTpdeductsuccessRequest 从 sync.Pool 获取 TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest
func GetTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest() *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest {
	return poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest.Get().(*TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest)
}

// ReleaseTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest 将 TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest 放入 sync.Pool
func ReleaseTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest(v *TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest) {
	v.Reset()
	poolTaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest.Put(v)
}
