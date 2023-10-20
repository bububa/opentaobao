package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelrategetAPIRequest 酒店产品库rate查询 API请求
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
type TaobaoxhotelrategetAPIRequest struct {
	model.Params
	// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
	_vendor string
	// 卖家自己系统的Code，简称RateCode
	_rateplanCode string
	// 卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合
	_outRid string
	// gid酒店商品id
	_gid int64
	// 酒店RPID
	_rpid int64
	// RateID
	_rateId int64
}

// NewTaobaoxhotelrategetRequest 初始化TaobaoxhotelrategetAPIRequest对象
func NewTaobaoxhotelrategetRequest() *TaobaoxhotelrategetAPIRequest {
	return &TaobaoxhotelrategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelrategetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelrategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelrategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 用于标示该宝贝的售卖渠道信息，允许同一个卖家酒店房型在淘宝系统发布多个售卖渠道的宝贝的价格。
func (r *TaobaoxhotelrategetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelrategetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRateplanCode is RateplanCode Setter
// 卖家自己系统的Code，简称RateCode
func (r *TaobaoxhotelrategetAPIRequest) SetRateplanCode(_rateplanCode string) error {
	r._rateplanCode = _rateplanCode
	r.Set("rateplan_code", _rateplanCode)
	return nil
}

// GetRateplanCode RateplanCode Getter
func (r TaobaoxhotelrategetAPIRequest) GetRateplanCode() string {
	return r._rateplanCode
}

// SetOutRid is OutRid Setter
// 卖家房型ID, 这是卖家自己系统中的房型ID 注意：需要按照规则组合
func (r *TaobaoxhotelrategetAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelrategetAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetGid is Gid Setter
// gid酒店商品id
func (r *TaobaoxhotelrategetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoxhotelrategetAPIRequest) GetGid() int64 {
	return r._gid
}

// SetRpid is Rpid Setter
// 酒店RPID
func (r *TaobaoxhotelrategetAPIRequest) SetRpid(_rpid int64) error {
	r._rpid = _rpid
	r.Set("rpid", _rpid)
	return nil
}

// GetRpid Rpid Getter
func (r TaobaoxhotelrategetAPIRequest) GetRpid() int64 {
	return r._rpid
}

// SetRateId is RateId Setter
// RateID
func (r *TaobaoxhotelrategetAPIRequest) SetRateId(_rateId int64) error {
	r._rateId = _rateId
	r.Set("rate_id", _rateId)
	return nil
}

// GetRateId RateId Getter
func (r TaobaoxhotelrategetAPIRequest) GetRateId() int64 {
	return r._rateId
}
