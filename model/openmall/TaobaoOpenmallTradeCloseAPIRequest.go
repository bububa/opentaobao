package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeCloseAPIRequest 关闭订单 API请求
// taobao.openmall.trade.close
//
// 关闭订单
type TaobaoOpenmallTradeCloseAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 关单原因
	_reason string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoOpenmallTradeCloseRequest 初始化TaobaoOpenmallTradeCloseAPIRequest对象
func NewTaobaoOpenmallTradeCloseRequest() *TaobaoOpenmallTradeCloseAPIRequest {
	return &TaobaoOpenmallTradeCloseAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTradeCloseAPIRequest) Reset() {
	r._distributor = ""
	r._reason = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeCloseAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeCloseAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeCloseAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetReason is Reason Setter
// 关单原因
func (r *TaobaoOpenmallTradeCloseAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoOpenmallTradeCloseAPIRequest) GetReason() string {
	return r._reason
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeCloseAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTradeCloseAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenmallTradeCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTradeCloseRequest()
	},
}

// GetTaobaoOpenmallTradeCloseRequest 从 sync.Pool 获取 TaobaoOpenmallTradeCloseAPIRequest
func GetTaobaoOpenmallTradeCloseAPIRequest() *TaobaoOpenmallTradeCloseAPIRequest {
	return poolTaobaoOpenmallTradeCloseAPIRequest.Get().(*TaobaoOpenmallTradeCloseAPIRequest)
}

// ReleaseTaobaoOpenmallTradeCloseAPIRequest 将 TaobaoOpenmallTradeCloseAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTradeCloseAPIRequest(v *TaobaoOpenmallTradeCloseAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTradeCloseAPIRequest.Put(v)
}
