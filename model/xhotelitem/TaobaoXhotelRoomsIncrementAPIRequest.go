package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomsincrementAPIRequest 房型库存推送接口（批量增量） API请求
// taobao.xhotel.rooms.increment
//
// Room库存增量更新接口，用户仅需要更新ROOM库存中发生变化的库存日历即可。
type TaobaoxhotelroomsincrementAPIRequest struct {
	model.Params
	// 批量全量推送房型共享库存,一次最多修改30个房型。json encode。示例：[{"out_rid":"hotel1_roomtype22","vendor":"","allotment_start_Time":"","allotment_end_time":"","superbook_start_time":"","superbook_end_time":"","roomQuota":[{"date":2010-01-28,"quota":10,"al_quota":2,"sp_quota":3}]}] 其中al_quota为保留房库存，sp_quota为超预定库存，quota为物理库存。al_quota和sp_quota需要向运营申请权限才可维护。allotment_start_Time和allotment_end_time为保留房库存开始和截止时间；superbook_start_time和superbook_end_time为超预定库存开始和截止时间。
	_roomQuotaMap string
}

// NewTaobaoxhotelroomsincrementRequest 初始化TaobaoxhotelroomsincrementAPIRequest对象
func NewTaobaoxhotelroomsincrementRequest() *TaobaoxhotelroomsincrementAPIRequest {
	return &TaobaoxhotelroomsincrementAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelroomsincrementAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rooms.increment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelroomsincrementAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelroomsincrementAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoomQuotaMap is RoomQuotaMap Setter
// 批量全量推送房型共享库存,一次最多修改30个房型。json encode。示例：[{&#34;out_rid&#34;:&#34;hotel1_roomtype22&#34;,&#34;vendor&#34;:&#34;&#34;,&#34;allotment_start_Time&#34;:&#34;&#34;,&#34;allotment_end_time&#34;:&#34;&#34;,&#34;superbook_start_time&#34;:&#34;&#34;,&#34;superbook_end_time&#34;:&#34;&#34;,&#34;roomQuota&#34;:[{&#34;date&#34;:2010-01-28,&#34;quota&#34;:10,&#34;al_quota&#34;:2,&#34;sp_quota&#34;:3}]}] 其中al_quota为保留房库存，sp_quota为超预定库存，quota为物理库存。al_quota和sp_quota需要向运营申请权限才可维护。allotment_start_Time和allotment_end_time为保留房库存开始和截止时间；superbook_start_time和superbook_end_time为超预定库存开始和截止时间。
func (r *TaobaoxhotelroomsincrementAPIRequest) SetRoomQuotaMap(_roomQuotaMap string) error {
	r._roomQuotaMap = _roomQuotaMap
	r.Set("room_quota_map", _roomQuotaMap)
	return nil
}

// GetRoomQuotaMap RoomQuotaMap Getter
func (r TaobaoxhotelroomsincrementAPIRequest) GetRoomQuotaMap() string {
	return r._roomQuotaMap
}
