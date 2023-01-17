package alsc

// MissionQuery 结构体
type MissionQuery struct {
	// 位置信息
	LocationInfos []LocationInfo `json:"location_infos,omitempty" xml:"location_infos>location_info,omitempty"`
	// 任务id列表
	MissionIds []int64 `json:"mission_ids,omitempty" xml:"mission_ids>int64,omitempty"`
	// 任务xid
	MissionXIds []string `json:"mission_x_ids,omitempty" xml:"mission_x_ids>string,omitempty"`
	//  控制是否展示查询任务响应结果里的某个字段
	RequestDataEnums []string `json:"request_data_enums,omitempty" xml:"request_data_enums>string,omitempty"`
	// 业务场景
	BizScence string `json:"biz_scence,omitempty" xml:"biz_scence,omitempty"`
	// 设备信息
	DeviceInfo *DeviceInfo `json:"device_info,omitempty" xml:"device_info,omitempty"`
	// 任务集id
	MissionCollectionId int64 `json:"mission_collection_id,omitempty" xml:"mission_collection_id,omitempty"`
	// 任务id
	MissionId int64 `json:"mission_id,omitempty" xml:"mission_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
	// 领取任务是否忽略投放
	IgnoreRecvTaskAladdin bool `json:"ignore_recv_task_aladdin,omitempty" xml:"ignore_recv_task_aladdin,omitempty"`
	// 是否走自定义排序
	IsSelfSort bool `json:"is_self_sort,omitempty" xml:"is_self_sort,omitempty"`
	// 本次查询，只查已领取的任务
	QueryReceivedFlag bool `json:"query_received_flag,omitempty" xml:"query_received_flag,omitempty"`
	// 是否展示提示信息（小红点）
	ShowRemind bool `json:"show_remind,omitempty" xml:"show_remind,omitempty"`
	// 是否单独展示组合任务的子任务列表
	ShowSublistAlone bool `json:"show_sublist_alone,omitempty" xml:"show_sublist_alone,omitempty"`
}
