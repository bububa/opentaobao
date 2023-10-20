package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTraceSearchAPIRequest 获取Openmall订单物流流转信息 API请求
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
type TaobaoOpenmallTraceSearchAPIRequest struct {
	model.Params
	// 签约支付宝代扣的账号
	_distributor string
	// 淘宝订单编号
	_tid int64
}

// NewTaobaoOpenmallTraceSearchRequest 初始化TaobaoOpenmallTraceSearchAPIRequest对象
func NewTaobaoOpenmallTraceSearchRequest() *TaobaoOpenmallTraceSearchAPIRequest {
	return &TaobaoOpenmallTraceSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTraceSearchAPIRequest) Reset() {
	r._distributor = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTraceSearchAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trace.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTraceSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTraceSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 签约支付宝代扣的账号
func (r *TaobaoOpenmallTraceSearchAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTraceSearchAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝订单编号
func (r *TaobaoOpenmallTraceSearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTraceSearchAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenmallTraceSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTraceSearchRequest()
	},
}

// GetTaobaoOpenmallTraceSearchRequest 从 sync.Pool 获取 TaobaoOpenmallTraceSearchAPIRequest
func GetTaobaoOpenmallTraceSearchAPIRequest() *TaobaoOpenmallTraceSearchAPIRequest {
	return poolTaobaoOpenmallTraceSearchAPIRequest.Get().(*TaobaoOpenmallTraceSearchAPIRequest)
}

// ReleaseTaobaoOpenmallTraceSearchAPIRequest 将 TaobaoOpenmallTraceSearchAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTraceSearchAPIRequest(v *TaobaoOpenmallTraceSearchAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTraceSearchAPIRequest.Put(v)
}
