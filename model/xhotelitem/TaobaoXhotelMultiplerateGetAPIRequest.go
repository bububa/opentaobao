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
	// 卖家的房价code
	_ratePlanCode string
	// 卖家的房型code
	_outRid string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 连住天数(范围1~10)
	_nod int64
	// 入住人数(范围1~10)
	_nop int64
	// 废弃，使用rate_plan_code
	_ratePlanId int64
	// 废弃，使用out_rid
	_gid int64
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
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRatePlanCode is RatePlanCode Setter
// 卖家的房价code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetRatePlanCode(_ratePlanCode string) error {
	r._ratePlanCode = _ratePlanCode
	r.Set("rate_plan_code", _ratePlanCode)
	return nil
}

// GetRatePlanCode RatePlanCode Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetRatePlanCode() string {
	return r._ratePlanCode
}

// SetOutRid is OutRid Setter
// 卖家的房型code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetNod is Nod Setter
// 连住天数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetNod(_nod int64) error {
	r._nod = _nod
	r.Set("nod", _nod)
	return nil
}

// GetNod Nod Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetNod() int64 {
	return r._nod
}

// SetNop is Nop Setter
// 入住人数(范围1~10)
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetNop(_nop int64) error {
	r._nop = _nop
	r.Set("nop", _nop)
	return nil
}

// GetNop Nop Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetNop() int64 {
	return r._nop
}

// SetRatePlanId is RatePlanId Setter
// 废弃，使用rate_plan_code
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetRatePlanId(_ratePlanId int64) error {
	r._ratePlanId = _ratePlanId
	r.Set("rate_plan_id", _ratePlanId)
	return nil
}

// GetRatePlanId RatePlanId Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetRatePlanId() int64 {
	return r._ratePlanId
}

// SetGid is Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelMultiplerateGetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoXhotelMultiplerateGetAPIRequest) GetGid() int64 {
	return r._gid
}
