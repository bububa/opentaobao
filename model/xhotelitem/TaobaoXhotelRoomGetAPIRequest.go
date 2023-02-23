package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomGetAPIRequest room实体查询 API请求
// taobao.xhotel.room.get
//
// 此接口用于查询一个商品，根据传入的gid查询商品信息。卖家只能查询自己的商品。
type TaobaoXhotelRoomGetAPIRequest struct {
	model.Params
	// 卖家渠道 如果gid为空，那么out_rid和vendor都不能为空。 支持通过gid或者通过out_rid和vendor来获取商品
	_vendor string
	// 外部房型id 如果gid为空，那么out_rid和vendor都不能为空 支持通过gid或者通过out_rid和vendor来获取商品
	_outRid string
	// 废弃
	_gid int64
}

// NewTaobaoXhotelRoomGetRequest 初始化TaobaoXhotelRoomGetAPIRequest对象
func NewTaobaoXhotelRoomGetRequest() *TaobaoXhotelRoomGetAPIRequest {
	return &TaobaoXhotelRoomGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.room.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// 卖家渠道 如果gid为空，那么out_rid和vendor都不能为空。 支持通过gid或者通过out_rid和vendor来获取商品
func (r *TaobaoXhotelRoomGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRoomGetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutRid is OutRid Setter
// 外部房型id 如果gid为空，那么out_rid和vendor都不能为空 支持通过gid或者通过out_rid和vendor来获取商品
func (r *TaobaoXhotelRoomGetAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoXhotelRoomGetAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetGid is Gid Setter
// 废弃
func (r *TaobaoXhotelRoomGetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoXhotelRoomGetAPIRequest) GetGid() int64 {
	return r._gid
}
