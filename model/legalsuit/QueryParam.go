package legalsuit

// QueryParam 结构体
type QueryParam struct {
	// 滑动标识（up:上滑  down：下滑）
	RollFlag string `json:"roll_flag,omitempty" xml:"roll_flag,omitempty"`
	// 业务id
	BusiId string `json:"busi_id,omitempty" xml:"busi_id,omitempty"`
	// 关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 操作人工号
	OperateWorkNo string `json:"operate_work_no,omitempty" xml:"operate_work_no,omitempty"`
	// 系统标识
	InputSystemCode string `json:"input_system_code,omitempty" xml:"input_system_code,omitempty"`
	// 场景id
	SceneId int64 `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	// 口径id
	StandpointId int64 `json:"standpoint_id,omitempty" xml:"standpoint_id,omitempty"`
}
