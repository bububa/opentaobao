package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeConfirmAPIRequest 确认收货 API请求
// taobao.openmall.trade.confirm
//
// 确认订单收货
type TaobaoOpenmallTradeConfirmAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoOpenmallTradeConfirmRequest 初始化TaobaoOpenmallTradeConfirmAPIRequest对象
func NewTaobaoOpenmallTradeConfirmRequest() *TaobaoOpenmallTradeConfirmAPIRequest {
	return &TaobaoOpenmallTradeConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeConfirmAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeConfirmAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeConfirmAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTradeConfirmAPIRequest) GetTid() int64 {
	return r._tid
}
