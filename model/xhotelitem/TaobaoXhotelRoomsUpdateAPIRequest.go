package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomsUpdateAPIRequest 房型共享库存推送接口（批量全量） API请求
// taobao.xhotel.rooms.update
//
// 此接口用于更新多个集市酒店商品房态信息，根据传入的gids更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在淘宝集市酒店系统中不存在，则会返回错误提示。是全量更新，非增量，会把之前的房态进行覆盖。
type TaobaoXhotelRoomsUpdateAPIRequest struct {
	model.Params
	// 批量全量推送房型共享库存,一次最多修改30个房型。json encode。示例：[{"out_rid":"hotel1_roomtype22","vendor":"","allotment_start_Time":"","allotment_end_time":"","superbook_start_time":"","superbook_end_time":"","roomQuota":[{"date":2010-01-28,"quota":10,"al_quota":2,"sp_quota":3}]}] 其中al_quota为保留房库存，sp_quota为超预定库存，quota为物理库存。al_quota和sp_quota需要向运营申请权限才可维护。allotment_start_Time和allotment_end_time为保留房库存开始和截止时间；superbook_start_time和superbook_end_time为超预定库存开始和截止时间。
	_roomQuotaMap string
}

// NewTaobaoXhotelRoomsUpdateRequest 初始化TaobaoXhotelRoomsUpdateAPIRequest对象
func NewTaobaoXhotelRoomsUpdateRequest() *TaobaoXhotelRoomsUpdateAPIRequest {
	return &TaobaoXhotelRoomsUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomsUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rooms.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomsUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomsUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRoomQuotaMap is RoomQuotaMap Setter
// 批量全量推送房型共享库存,一次最多修改30个房型。json encode。示例：[{&#34;out_rid&#34;:&#34;hotel1_roomtype22&#34;,&#34;vendor&#34;:&#34;&#34;,&#34;allotment_start_Time&#34;:&#34;&#34;,&#34;allotment_end_time&#34;:&#34;&#34;,&#34;superbook_start_time&#34;:&#34;&#34;,&#34;superbook_end_time&#34;:&#34;&#34;,&#34;roomQuota&#34;:[{&#34;date&#34;:2010-01-28,&#34;quota&#34;:10,&#34;al_quota&#34;:2,&#34;sp_quota&#34;:3}]}] 其中al_quota为保留房库存，sp_quota为超预定库存，quota为物理库存。al_quota和sp_quota需要向运营申请权限才可维护。allotment_start_Time和allotment_end_time为保留房库存开始和截止时间；superbook_start_time和superbook_end_time为超预定库存开始和截止时间。
func (r *TaobaoXhotelRoomsUpdateAPIRequest) SetRoomQuotaMap(_roomQuotaMap string) error {
	r._roomQuotaMap = _roomQuotaMap
	r.Set("room_quota_map", _roomQuotaMap)
	return nil
}

// GetRoomQuotaMap RoomQuotaMap Getter
func (r TaobaoXhotelRoomsUpdateAPIRequest) GetRoomQuotaMap() string {
	return r._roomQuotaMap
}
