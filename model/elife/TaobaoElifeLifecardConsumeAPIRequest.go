package elife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardConsumeAPIRequest 品牌惠卡券核销 API请求
// taobao.elife.lifecard.consume
//
// 用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作
type TaobaoElifeLifecardConsumeAPIRequest struct {
	model.Params
	// 交易请求参数
	_consumeRequest *ConsumeRequest
}

// NewTaobaoElifeLifecardConsumeRequest 初始化TaobaoElifeLifecardConsumeAPIRequest对象
func NewTaobaoElifeLifecardConsumeRequest() *TaobaoElifeLifecardConsumeAPIRequest {
	return &TaobaoElifeLifecardConsumeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoElifeLifecardConsumeAPIRequest) Reset() {
	r._consumeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardConsumeAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.consume"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardConsumeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoElifeLifecardConsumeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConsumeRequest is ConsumeRequest Setter
// 交易请求参数
func (r *TaobaoElifeLifecardConsumeAPIRequest) SetConsumeRequest(_consumeRequest *ConsumeRequest) error {
	r._consumeRequest = _consumeRequest
	r.Set("consume_request", _consumeRequest)
	return nil
}

// GetConsumeRequest ConsumeRequest Getter
func (r TaobaoElifeLifecardConsumeAPIRequest) GetConsumeRequest() *ConsumeRequest {
	return r._consumeRequest
}

var poolTaobaoElifeLifecardConsumeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoElifeLifecardConsumeRequest()
	},
}

// GetTaobaoElifeLifecardConsumeRequest 从 sync.Pool 获取 TaobaoElifeLifecardConsumeAPIRequest
func GetTaobaoElifeLifecardConsumeAPIRequest() *TaobaoElifeLifecardConsumeAPIRequest {
	return poolTaobaoElifeLifecardConsumeAPIRequest.Get().(*TaobaoElifeLifecardConsumeAPIRequest)
}

// ReleaseTaobaoElifeLifecardConsumeAPIRequest 将 TaobaoElifeLifecardConsumeAPIRequest 放入 sync.Pool
func ReleaseTaobaoElifeLifecardConsumeAPIRequest(v *TaobaoElifeLifecardConsumeAPIRequest) {
	v.Reset()
	poolTaobaoElifeLifecardConsumeAPIRequest.Put(v)
}
