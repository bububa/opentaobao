package openmall

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradegetAPIRequest 查询订单详情 API请求
// taobao.openmall.trade.get
//
// 查询订单详情
type TaobaoopenmalltradegetAPIRequest struct {
	model.Params
	// 分销者信息
	_distributor string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoopenmalltradegetRequest 初始化TaobaoopenmalltradegetAPIRequest对象
func NewTaobaoopenmalltradegetRequest() *TaobaoopenmalltradegetAPIRequest {
	return &TaobaoopenmalltradegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenmalltradegetAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenmalltradegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenmalltradegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 分销者信息
func (r *TaobaoopenmalltradegetAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoopenmalltradegetAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoopenmalltradegetAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoopenmalltradegetAPIRequest) GetTid() int64 {
	return r._tid
}
