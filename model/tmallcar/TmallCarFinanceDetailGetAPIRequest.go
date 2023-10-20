package tmallcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFinanceDetailGetAPIRequest 查询汽车金融订单信息 API请求
// tmall.car.finance.detail.get
//
// 查询汽车金融订单信息
type TmallCarFinanceDetailGetAPIRequest struct {
	model.Params
	// 查询参数
	_param0 *FinanceDetailQueryReq
}

// NewTmallCarFinanceDetailGetRequest 初始化TmallCarFinanceDetailGetAPIRequest对象
func NewTmallCarFinanceDetailGetRequest() *TmallCarFinanceDetailGetAPIRequest {
	return &TmallCarFinanceDetailGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCarFinanceDetailGetAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarFinanceDetailGetAPIRequest) GetApiMethodName() string {
	return "tmall.car.finance.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarFinanceDetailGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarFinanceDetailGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 查询参数
func (r *TmallCarFinanceDetailGetAPIRequest) SetParam0(_param0 *FinanceDetailQueryReq) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TmallCarFinanceDetailGetAPIRequest) GetParam0() *FinanceDetailQueryReq {
	return r._param0
}

var poolTmallCarFinanceDetailGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCarFinanceDetailGetRequest()
	},
}

// GetTmallCarFinanceDetailGetRequest 从 sync.Pool 获取 TmallCarFinanceDetailGetAPIRequest
func GetTmallCarFinanceDetailGetAPIRequest() *TmallCarFinanceDetailGetAPIRequest {
	return poolTmallCarFinanceDetailGetAPIRequest.Get().(*TmallCarFinanceDetailGetAPIRequest)
}

// ReleaseTmallCarFinanceDetailGetAPIRequest 将 TmallCarFinanceDetailGetAPIRequest 放入 sync.Pool
func ReleaseTmallCarFinanceDetailGetAPIRequest(v *TmallCarFinanceDetailGetAPIRequest) {
	v.Reset()
	poolTmallCarFinanceDetailGetAPIRequest.Put(v)
}
