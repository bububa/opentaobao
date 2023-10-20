package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomtypedeletepublicAPIRequest 商家删除房型数据接口 API请求
// taobao.xhotel.roomtype.delete.public
//
// 房型删除TOP接口
type TaobaoxhotelroomtypedeletepublicAPIRequest struct {
	model.Params
	// vendor
	_vendor string
	// 外部房型ID
	_outerRid string
	// 具体操作人，比如酒店帐号、小二名称等
	_operator string
	// 房型rid ，传参方式：rid    或者   outer_id+vendor
	_rid int64
}

// NewTaobaoxhotelroomtypedeletepublicRequest 初始化TaobaoxhotelroomtypedeletepublicAPIRequest对象
func NewTaobaoxhotelroomtypedeletepublicRequest() *TaobaoxhotelroomtypedeletepublicAPIRequest {
	return &TaobaoxhotelroomtypedeletepublicAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.delete.public"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVendor is Vendor Setter
// vendor
func (r *TaobaoxhotelroomtypedeletepublicAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOuterRid is OuterRid Setter
// 外部房型ID
func (r *TaobaoxhotelroomtypedeletepublicAPIRequest) SetOuterRid(_outerRid string) error {
	r._outerRid = _outerRid
	r.Set("outer_rid", _outerRid)
	return nil
}

// GetOuterRid OuterRid Getter
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetOuterRid() string {
	return r._outerRid
}

// SetOperator is Operator Setter
// 具体操作人，比如酒店帐号、小二名称等
func (r *TaobaoxhotelroomtypedeletepublicAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetOperator() string {
	return r._operator
}

// SetRid is Rid Setter
// 房型rid ，传参方式：rid    或者   outer_id+vendor
func (r *TaobaoxhotelroomtypedeletepublicAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoxhotelroomtypedeletepublicAPIRequest) GetRid() int64 {
	return r._rid
}
