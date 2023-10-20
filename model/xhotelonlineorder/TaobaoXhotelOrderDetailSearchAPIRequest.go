package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelorderdetailsearchAPIRequest 订单详情查询 API请求
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
type TaobaoxhotelorderdetailsearchAPIRequest struct {
	model.Params
	// 外部订单号
	_outOid string
	// 外部订单号
	_tid int64
}

// NewTaobaoxhotelorderdetailsearchRequest 初始化TaobaoxhotelorderdetailsearchAPIRequest对象
func NewTaobaoxhotelorderdetailsearchRequest() *TaobaoxhotelorderdetailsearchAPIRequest {
	return &TaobaoxhotelorderdetailsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelorderdetailsearchAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.detail.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelorderdetailsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelorderdetailsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutOid is OutOid Setter
// 外部订单号
func (r *TaobaoxhotelorderdetailsearchAPIRequest) SetOutOid(_outOid string) error {
	r._outOid = _outOid
	r.Set("out_oid", _outOid)
	return nil
}

// GetOutOid OutOid Getter
func (r TaobaoxhotelorderdetailsearchAPIRequest) GetOutOid() string {
	return r._outOid
}

// SetTid is Tid Setter
// 外部订单号
func (r *TaobaoxhotelorderdetailsearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoxhotelorderdetailsearchAPIRequest) GetTid() int64 {
	return r._tid
}
