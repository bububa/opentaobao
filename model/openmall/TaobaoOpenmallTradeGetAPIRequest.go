package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeGetAPIRequest 查询订单详情 API请求
// taobao.openmall.trade.get
//
// 查询订单详情
type TaobaoOpenmallTradeGetAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoOpenmallTradeGetRequest 初始化TaobaoOpenmallTradeGetAPIRequest对象
func NewTaobaoOpenmallTradeGetRequest() *TaobaoOpenmallTradeGetAPIRequest {
	return &TaobaoOpenmallTradeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTradeGetAPIRequest) Reset() {
	r._distributor = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTradeGetAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenmallTradeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTradeGetRequest()
	},
}

// GetTaobaoOpenmallTradeGetRequest 从 sync.Pool 获取 TaobaoOpenmallTradeGetAPIRequest
func GetTaobaoOpenmallTradeGetAPIRequest() *TaobaoOpenmallTradeGetAPIRequest {
	return poolTaobaoOpenmallTradeGetAPIRequest.Get().(*TaobaoOpenmallTradeGetAPIRequest)
}

// ReleaseTaobaoOpenmallTradeGetAPIRequest 将 TaobaoOpenmallTradeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTradeGetAPIRequest(v *TaobaoOpenmallTradeGetAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTradeGetAPIRequest.Put(v)
}
