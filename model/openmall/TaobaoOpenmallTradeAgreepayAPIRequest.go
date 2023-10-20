package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeAgreepayAPIRequest openmall订单支付 API请求
// taobao.openmall.trade.agreepay
//
// openmall订单支付
type TaobaoOpenmallTradeAgreepayAPIRequest struct {
	model.Params
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	_distributor string
	// 淘宝交易单号
	_tid int64
}

// NewTaobaoOpenmallTradeAgreepayRequest 初始化TaobaoOpenmallTradeAgreepayAPIRequest对象
func NewTaobaoOpenmallTradeAgreepayRequest() *TaobaoOpenmallTradeAgreepayAPIRequest {
	return &TaobaoOpenmallTradeAgreepayAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTradeAgreepayAPIRequest) Reset() {
	r._distributor = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.agreepay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
func (r *TaobaoOpenmallTradeAgreepayAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝交易单号
func (r *TaobaoOpenmallTradeAgreepayAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTradeAgreepayAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenmallTradeAgreepayAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTradeAgreepayRequest()
	},
}

// GetTaobaoOpenmallTradeAgreepayRequest 从 sync.Pool 获取 TaobaoOpenmallTradeAgreepayAPIRequest
func GetTaobaoOpenmallTradeAgreepayAPIRequest() *TaobaoOpenmallTradeAgreepayAPIRequest {
	return poolTaobaoOpenmallTradeAgreepayAPIRequest.Get().(*TaobaoOpenmallTradeAgreepayAPIRequest)
}

// ReleaseTaobaoOpenmallTradeAgreepayAPIRequest 将 TaobaoOpenmallTradeAgreepayAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTradeAgreepayAPIRequest(v *TaobaoOpenmallTradeAgreepayAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTradeAgreepayAPIRequest.Put(v)
}
