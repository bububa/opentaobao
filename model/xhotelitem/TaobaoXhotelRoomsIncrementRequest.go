package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型库存推送接口（批量增量） APIRequest
taobao.xhotel.rooms.increment

Room库存增量更新接口，用户仅需要更新ROOM库存中发生变化的库存日历即可。
*/
type TaobaoXhotelRoomsIncrementRequest struct {
    model.Params

    // 批量全量推送房型共享库存,一次最多修改30个房型。json encode。示例：[{"out_rid":"hotel1_roomtype22","vendor":"","allotment_start_Time":"","allotment_end_time":"","superbook_start_time":"","superbook_end_time":"","roomQuota":[{"date":2010-01-28,"quota":10,"al_quota":2,"sp_quota":3}]}] 其中al_quota为保留房库存，sp_quota为超预定库存，quota为物理库存。al_quota和sp_quota需要向运营申请权限才可维护。allotment_start_Time和allotment_end_time为保留房库存开始和截止时间；superbook_start_time和superbook_end_time为超预定库存开始和截止时间。
    roomQuotaMap   string 

}

func NewTaobaoXhotelRoomsIncrementRequest() *TaobaoXhotelRoomsIncrementRequest{
    return &TaobaoXhotelRoomsIncrementRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelRoomsIncrementRequest) GetApiMethodName() string {
    return "taobao.xhotel.rooms.increment"
}

func (r TaobaoXhotelRoomsIncrementRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelRoomsIncrementRequest) SetRoomQuotaMap(roomQuotaMap string) error {
    r.roomQuotaMap = roomQuotaMap
    r.Set("room_quota_map", roomQuotaMap)
    return nil
}

func (r TaobaoXhotelRoomsIncrementRequest) GetRoomQuotaMap() string {
    return r.roomQuotaMap
}

