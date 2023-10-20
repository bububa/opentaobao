package alsc

import (
	"sync"
)

// MissionReceiveQuery 结构体
type MissionReceiveQuery struct {
	// 位置信息
	LocationInfos []LocationInfo `json:"location_infos,omitempty" xml:"location_infos>location_info,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 任务xid
	MissionXId string `json:"mission_x_id,omitempty" xml:"mission_x_id,omitempty"`
	// 排序分组id，目前井冈山业务使用，任务分组
	SortGroupId string `json:"sort_group_id,omitempty" xml:"sort_group_id,omitempty"`
	// 设备号
	DeviceInfo *DeviceInfo `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 任务集id
	MissionCollectionId int64 `json:"mission_collection_id,omitempty" xml:"mission_collection_id,omitempty"`
	// 任务id
	MissionDefId int64 `json:"mission_def_id,omitempty" xml:"mission_def_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolMissionReceiveQuery = sync.Pool{
	New: func() any {
		return new(MissionReceiveQuery)
	},
}

// GetMissionReceiveQuery() 从对象池中获取MissionReceiveQuery
func GetMissionReceiveQuery() *MissionReceiveQuery {
	return poolMissionReceiveQuery.Get().(*MissionReceiveQuery)
}

// ReleaseMissionReceiveQuery 释放MissionReceiveQuery
func ReleaseMissionReceiveQuery(v *MissionReceiveQuery) {
	v.LocationInfos = v.LocationInfos[:0]
	v.BizScene = ""
	v.MissionXId = ""
	v.SortGroupId = ""
	v.DeviceInfo = nil
	v.MissionCollectionId = 0
	v.MissionDefId = 0
	v.UserInfo = nil
	poolMissionReceiveQuery.Put(v)
}
