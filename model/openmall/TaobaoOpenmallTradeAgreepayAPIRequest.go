package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradeagreepayAPIRequest openmall订单支付 API请求
// taobao.openmall.trade.agreepay
//
// openmall订单支付
type TaobaoopenmalltradeagreepayAPIRequest struct {
	model.Params
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	_distributor string
	// 淘宝交易单号
	_tid int64
}

// NewTaobaoopenmalltradeagreepayRequest 初始化TaobaoopenmalltradeagreepayAPIRequest对象
func NewTaobaoopenmalltradeagreepayRequest() *TaobaoopenmalltradeagreepayAPIRequest {
	return &TaobaoopenmalltradeagreepayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradeagreepayAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.agreepay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradeagreepayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradeagreepayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
func (r *TaobaoopenmalltradeagreepayAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltradeagreepayAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝交易单号
func (r *TaobaoopenmalltradeagreepayAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenmalltradeagreepayAPIRequest) GetTid() int64 {
	return r._tid
}
