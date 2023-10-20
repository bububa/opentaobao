package trade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoRefundRefundactionsJudgementAPIRequest 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API请求
// cainiao.refund.refundactions.judgement
//
// 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
type CainiaoRefundRefundactionsJudgementAPIRequest struct {
	model.Params
	// 操作请求
	_param0 *OrderRefundOperationJudgementReq
}

// NewCainiaoRefundRefundactionsJudgementRequest 初始化CainiaoRefundRefundactionsJudgementAPIRequest对象
func NewCainiaoRefundRefundactionsJudgementRequest() *CainiaoRefundRefundactionsJudgementAPIRequest {
	return &CainiaoRefundRefundactionsJudgementAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoRefundRefundactionsJudgementAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoRefundRefundactionsJudgementAPIRequest) GetApiMethodName() string {
	return "cainiao.refund.refundactions.judgement"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoRefundRefundactionsJudgementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoRefundRefundactionsJudgementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 操作请求
func (r *CainiaoRefundRefundactionsJudgementAPIRequest) SetParam0(_param0 *OrderRefundOperationJudgementReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r CainiaoRefundRefundactionsJudgementAPIRequest) GetParam0() *OrderRefundOperationJudgementReq {
	return r._param0
}

var poolCainiaoRefundRefundactionsJudgementAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoRefundRefundactionsJudgementRequest()
	},
}

// GetCainiaoRefundRefundactionsJudgementRequest 从 sync.Pool 获取 CainiaoRefundRefundactionsJudgementAPIRequest
func GetCainiaoRefundRefundactionsJudgementAPIRequest() *CainiaoRefundRefundactionsJudgementAPIRequest {
	return poolCainiaoRefundRefundactionsJudgementAPIRequest.Get().(*CainiaoRefundRefundactionsJudgementAPIRequest)
}

// ReleaseCainiaoRefundRefundactionsJudgementAPIRequest 将 CainiaoRefundRefundactionsJudgementAPIRequest 放入 sync.Pool
func ReleaseCainiaoRefundRefundactionsJudgementAPIRequest(v *CainiaoRefundRefundactionsJudgementAPIRequest) {
	v.Reset()
	poolCainiaoRefundRefundactionsJudgementAPIRequest.Put(v)
}
