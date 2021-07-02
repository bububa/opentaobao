package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateGetAPIRequest 复杂房价查询接口 API请求
// taobao.xhotel.multiplerate.get
//
// 查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
type TaobaoXhotelMultiplerateGetAPIRequest struct {
	model.Params
	// 连住天数(范围1~10)
	_nod int64
	// 入住人数(范围1~10)
	_nop int64
	// 卖家的房价code
	_ratePlanCode string
	// 废弃，使用rate_plan_code
	_ratePlanId int64
	// 卖家的房型code
	_outRid string
	// 废弃，使用out_rid
	_gid int64
	// 系统商，一般不填写，使用须申请
	_vendor string
}

// NewTaobaoXhotelMultiplerateGetRequest 初始化TaobaoXhotelMultiplerateGetAPIRequest对象
func NewTaobaoXhotelMultiplerateGetRequest() *TaobaoXhotelMultiplerateGetAPIRequest {
	return &TaobaoXhotelMultiplerateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.multiplerate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nod Setter
// 连住天数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetNod(_nod int64) error {
	r._nod = _nod
	r.Set("nod", _nod)
	return nil
}

// Get Nod Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetNod() int64 {
	return r._nod
}

// Set is Nop Setter
// 入住人数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetNop(_nop int64) error {
	r._nop = _nop
	r.Set("nop", _nop)
	return nil
}

// Get Nop Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetNop() int64 {
	return r._nop
}

// Set is RatePlanCode Setter
// 卖家的房价code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetRatePlanCode(_ratePlanCode string) error {
	r._ratePlanCode = _ratePlanCode
	r.Set("rate_plan_code", _ratePlanCode)
	return nil
}

// Get RatePlanCode Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetRatePlanCode() string {
	return r._ratePlanCode
}

// Set is RatePlanId Setter
// 废弃，使用rate_plan_code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetRatePlanId(_ratePlanId int64) error {
	r._ratePlanId = _ratePlanId
	r.Set("rate_plan_id", _ratePlanId)
	return nil
}

// Get RatePlanId Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetRatePlanId() int64 {
	return r._ratePlanId
}

// Set is OutRid Setter
// 卖家的房型code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// Get OutRid Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetOutRid() string {
	return r._outRid
}

// Set is Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// Get Gid Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetGid() int64 {
	return r._gid
}

// Set is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetVendor() string {
	return r._vendor
}
