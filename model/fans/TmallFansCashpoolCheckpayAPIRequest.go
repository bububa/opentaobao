package fans

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansCashpoolCheckpayAPIRequest 检查资金池付款状态 API请求
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
type TmallFansCashpoolCheckpayAPIRequest struct {
	model.Params
	// 资金池列表
	_cashPoolList []int64
}

// NewTmallFansCashpoolCheckpayRequest 初始化TmallFansCashpoolCheckpayAPIRequest对象
func NewTmallFansCashpoolCheckpayRequest() *TmallFansCashpoolCheckpayAPIRequest {
	return &TmallFansCashpoolCheckpayAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFansCashpoolCheckpayAPIRequest) Reset() {
	r._cashPoolList = r._cashPoolList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFansCashpoolCheckpayAPIRequest) GetApiMethodName() string {
	return "tmall.fans.cashpool.checkpay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFansCashpoolCheckpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFansCashpoolCheckpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCashPoolList is CashPoolList Setter
// 资金池列表
func (r *TmallFansCashpoolCheckpayAPIRequest) SetCashPoolList(_cashPoolList []int64) error {
	r._cashPoolList = _cashPoolList
	r.Set("cash_pool_list", _cashPoolList)
	return nil
}

// GetCashPoolList CashPoolList Getter
func (r TmallFansCashpoolCheckpayAPIRequest) GetCashPoolList() []int64 {
	return r._cashPoolList
}

var poolTmallFansCashpoolCheckpayAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFansCashpoolCheckpayRequest()
	},
}

// GetTmallFansCashpoolCheckpayRequest 从 sync.Pool 获取 TmallFansCashpoolCheckpayAPIRequest
func GetTmallFansCashpoolCheckpayAPIRequest() *TmallFansCashpoolCheckpayAPIRequest {
	return poolTmallFansCashpoolCheckpayAPIRequest.Get().(*TmallFansCashpoolCheckpayAPIRequest)
}

// ReleaseTmallFansCashpoolCheckpayAPIRequest 将 TmallFansCashpoolCheckpayAPIRequest 放入 sync.Pool
func ReleaseTmallFansCashpoolCheckpayAPIRequest(v *TmallFansCashpoolCheckpayAPIRequest) {
	v.Reset()
	poolTmallFansCashpoolCheckpayAPIRequest.Put(v)
}
