package idle

// IdleSpuRegister4TopDto 结构体
type IdleSpuRegister4TopDto struct {
	// 业务场景编码, 如YHB_3C、RECYCLE_PHONE，具体值由挂载spu文档给出
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
	// 操作类型，0新增，-1删除。当spu第一次挂载时，会进入1测试中状态。服务商联调通过后，需要再次挂载，actionType还传0，挂载信息状态会变成0已上线。
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// spuId 具体值由挂载spu文档给出
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
}
