package alsc

import (
	"sync"
)

// RewardReceiveQuery 结构体
type RewardReceiveQuery struct {
	// 位置信息
	LocationInfos []LocationInfo `json:"location_infos,omitempty" xml:"location_infos>location_info,omitempty"`
	// 业务场景
	BizScene string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 扩展信息
	Ext string `json:"ext,omitempty" xml:"ext,omitempty"`
	// 任务id
	MissionXId string `json:"mission_x_id,omitempty" xml:"mission_x_id,omitempty"`
	// 风控参数
	RiskParam string `json:"risk_param,omitempty" xml:"risk_param,omitempty"`
	// 设备信息
	DeviceInfo *DeviceInfo `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 任务集id
	MissionCollectionId int64 `json:"mission_collection_id,omitempty" xml:"mission_collection_id,omitempty"`
	// 任务id
	MissionDefId int64 `json:"mission_def_id,omitempty" xml:"mission_def_id,omitempty"`
	// 任务实例id 领取型任务必传
	MissionInstanceId int64 `json:"mission_instance_id,omitempty" xml:"mission_instance_id,omitempty"`
	// 阶段信息
	Stage *Stage `json:"stage,omitempty" xml:"stage,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolRewardReceiveQuery = sync.Pool{
	New: func() any {
		return new(RewardReceiveQuery)
	},
}

// GetRewardReceiveQuery() 从对象池中获取RewardReceiveQuery
func GetRewardReceiveQuery() *RewardReceiveQuery {
	return poolRewardReceiveQuery.Get().(*RewardReceiveQuery)
}

// ReleaseRewardReceiveQuery 释放RewardReceiveQuery
func ReleaseRewardReceiveQuery(v *RewardReceiveQuery) {
	v.LocationInfos = v.LocationInfos[:0]
	v.BizScene = ""
	v.Ext = ""
	v.MissionXId = ""
	v.RiskParam = ""
	v.DeviceInfo = nil
	v.MissionCollectionId = 0
	v.MissionDefId = 0
	v.MissionInstanceId = 0
	v.Stage = nil
	v.UserInfo = nil
	poolRewardReceiveQuery.Put(v)
}
