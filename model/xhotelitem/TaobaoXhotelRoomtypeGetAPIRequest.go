package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypegetAPIRequest 房型查询接口 API请求
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
type TaobaoxhotelroomtypegetAPIRequest struct {
	model.Params
	// 商家房型ID
	_outerId string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 废弃，使用商家房型ID
	_rid int64
}

// NewTaobaoxhotelroomtypegetRequest 初始化TaobaoxhotelroomtypegetAPIRequest对象
func NewTaobaoxhotelroomtypegetRequest() *TaobaoxhotelroomtypegetAPIRequest {
	return &TaobaoxhotelroomtypegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelroomtypegetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelroomtypegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelroomtypegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 商家房型ID
func (r *TaobaoxhotelroomtypegetAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelroomtypegetAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoxhotelroomtypegetAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelroomtypegetAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRid is Rid Setter
// 废弃，使用商家房型ID
func (r *TaobaoxhotelroomtypegetAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoxhotelroomtypegetAPIRequest) GetRid() int64 {
	return r._rid
}
