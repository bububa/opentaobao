package fans

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansCashpoolCreateAPIRequest 创建资金池 API请求
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
type TmallFansCashpoolCreateAPIRequest struct {
	model.Params
	// 创建资奖池输入对象
	_createCashPoolParamDo *CreateCashPoolParamDo
}

// NewTmallFansCashpoolCreateRequest 初始化TmallFansCashpoolCreateAPIRequest对象
func NewTmallFansCashpoolCreateRequest() *TmallFansCashpoolCreateAPIRequest {
	return &TmallFansCashpoolCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallFansCashpoolCreateAPIRequest) Reset() {
	r._createCashPoolParamDo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFansCashpoolCreateAPIRequest) GetApiMethodName() string {
	return "tmall.fans.cashpool.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFansCashpoolCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFansCashpoolCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateCashPoolParamDo is CreateCashPoolParamDo Setter
// 创建资奖池输入对象
func (r *TmallFansCashpoolCreateAPIRequest) SetCreateCashPoolParamDo(_createCashPoolParamDo *CreateCashPoolParamDo) error {
	r._createCashPoolParamDo = _createCashPoolParamDo
	r.Set("create_cash_pool_param_do", _createCashPoolParamDo)
	return nil
}

// GetCreateCashPoolParamDo CreateCashPoolParamDo Getter
func (r TmallFansCashpoolCreateAPIRequest) GetCreateCashPoolParamDo() *CreateCashPoolParamDo {
	return r._createCashPoolParamDo
}

var poolTmallFansCashpoolCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallFansCashpoolCreateRequest()
	},
}

// GetTmallFansCashpoolCreateRequest 从 sync.Pool 获取 TmallFansCashpoolCreateAPIRequest
func GetTmallFansCashpoolCreateAPIRequest() *TmallFansCashpoolCreateAPIRequest {
	return poolTmallFansCashpoolCreateAPIRequest.Get().(*TmallFansCashpoolCreateAPIRequest)
}

// ReleaseTmallFansCashpoolCreateAPIRequest 将 TmallFansCashpoolCreateAPIRequest 放入 sync.Pool
func ReleaseTmallFansCashpoolCreateAPIRequest(v *TmallFansCashpoolCreateAPIRequest) {
	v.Reset()
	poolTmallFansCashpoolCreateAPIRequest.Put(v)
}
