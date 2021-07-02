package openmall

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeGetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Distributor Setter
// 分销者信息
func (r *TaobaoOpenmallTradeGetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// Get Distributor Getter
func (r TaobaoOpenmallTradeGetAPIRequest) GetDistributor() string {
	return r._distributor
}

// Set is Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeGetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// Get Tid Getter
func (r TaobaoOpenmallTradeGetAPIRequest) GetTid() int64 {
	return r._tid
}
