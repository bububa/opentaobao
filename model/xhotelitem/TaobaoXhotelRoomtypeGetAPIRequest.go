package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeGetAPIRequest 房型查询接口 API请求
// taobao.xhotel.roomtype.get
//
// 房型查询房型查询接口返回结果增加date_confirm字段
type TaobaoXhotelRoomtypeGetAPIRequest struct {
	model.Params
	// 商家房型ID
	_outerId string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 废弃，使用商家房型ID
	_rid int64
}

// NewTaobaoXhotelRoomtypeGetRequest 初始化TaobaoXhotelRoomtypeGetAPIRequest对象
func NewTaobaoXhotelRoomtypeGetRequest() *TaobaoXhotelRoomtypeGetAPIRequest {
	return &TaobaoXhotelRoomtypeGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRoomtypeGetAPIRequest) Reset() {
	r._outerId = ""
	r._vendor = ""
	r._rid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomtypeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoXhotelRoomtypeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRoomtypeGetRequest()
	},
}

// GetTaobaoXhotelRoomtypeGetRequest 从 sync.Pool 获取 TaobaoXhotelRoomtypeGetAPIRequest
func GetTaobaoXhotelRoomtypeGetAPIRequest() *TaobaoXhotelRoomtypeGetAPIRequest {
	return poolTaobaoXhotelRoomtypeGetAPIRequest.Get().(*TaobaoXhotelRoomtypeGetAPIRequest)
}

// ReleaseTaobaoXhotelRoomtypeGetAPIRequest 将 TaobaoXhotelRoomtypeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRoomtypeGetAPIRequest(v *TaobaoXhotelRoomtypeGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRoomtypeGetAPIRequest.Put(v)
}
