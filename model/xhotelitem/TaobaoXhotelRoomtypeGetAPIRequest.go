package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeGetAPIRequest 房型查询接口 API请求
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
type TaobaoXhotelRoomtypeGetAPIRequest struct {
	model.Params
	// 废弃，使用商家房型ID
	_rid int64
	// 商家房型ID
	_outerId string
	// 系统商，一般不填写，使用须申请
	_vendor string
}

// NewTaobaoXhotelRoomtypeGetRequest 初始化TaobaoXhotelRoomtypeGetAPIRequest对象
func NewTaobaoXhotelRoomtypeGetRequest() *TaobaoXhotelRoomtypeGetAPIRequest {
	return &TaobaoXhotelRoomtypeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRid is Rid Setter
// 废弃，使用商家房型ID
func (r *TaobaoXhotelRoomtypeGetAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetRid() int64 {
	return r._rid
}

// SetOuterId is OuterId Setter
// 商家房型ID
func (r *TaobaoXhotelRoomtypeGetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomtypeGetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetVendor() string {
	return r._vendor
}
