package deliveryvoucher

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGameDeliveryvoucherEvaluateAPIRequest 卡券评价回传 API请求
// taobao.game.deliveryvoucher.evaluate
//
// 卡券ISV回传商品评价
type TaobaoGameDeliveryvoucherEvaluateAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *VoucherEvaluateRequest
}

// NewTaobaoGameDeliveryvoucherEvaluateRequest 初始化TaobaoGameDeliveryvoucherEvaluateAPIRequest对象
func NewTaobaoGameDeliveryvoucherEvaluateRequest() *TaobaoGameDeliveryvoucherEvaluateAPIRequest {
	return &TaobaoGameDeliveryvoucherEvaluateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGameDeliveryvoucherEvaluateAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetApiMethodName() string {
	return "taobao.game.deliveryvoucher.evaluate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 系统自动生成
func (r *TaobaoGameDeliveryvoucherEvaluateAPIRequest) SetParam0(_param0 *VoucherEvaluateRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoGameDeliveryvoucherEvaluateAPIRequest) GetParam0() *VoucherEvaluateRequest {
	return r._param0
}

var poolTaobaoGameDeliveryvoucherEvaluateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGameDeliveryvoucherEvaluateRequest()
	},
}

// GetTaobaoGameDeliveryvoucherEvaluateRequest 从 sync.Pool 获取 TaobaoGameDeliveryvoucherEvaluateAPIRequest
func GetTaobaoGameDeliveryvoucherEvaluateAPIRequest() *TaobaoGameDeliveryvoucherEvaluateAPIRequest {
	return poolTaobaoGameDeliveryvoucherEvaluateAPIRequest.Get().(*TaobaoGameDeliveryvoucherEvaluateAPIRequest)
}

// ReleaseTaobaoGameDeliveryvoucherEvaluateAPIRequest 将 TaobaoGameDeliveryvoucherEvaluateAPIRequest 放入 sync.Pool
func ReleaseTaobaoGameDeliveryvoucherEvaluateAPIRequest(v *TaobaoGameDeliveryvoucherEvaluateAPIRequest) {
	v.Reset()
	poolTaobaoGameDeliveryvoucherEvaluateAPIRequest.Put(v)
}
