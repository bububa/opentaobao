package alimember

import (
	"sync"
)

// PointOperateDto 结构体
type PointOperateDto struct {
	// 消费积分
	TotalPoint string `json:"total_point,omitempty" xml:"total_point,omitempty"`
	// 场景码，对接时业务分配
	SceneCode string `json:"scene_code,omitempty" xml:"scene_code,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 等级积分
	LevelPoint string `json:"level_point,omitempty" xml:"level_point,omitempty"`
	// 操作序列号，需要唯一
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 扩展字段
	ExtraInfo string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// 用户站点
	UserSite string `json:"user_site,omitempty" xml:"user_site,omitempty"`
	// 开放用户id
	OpenUserId string `json:"open_user_id,omitempty" xml:"open_user_id,omitempty"`
	// 过期时间，单位毫秒
	ExpireTime int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 操作类型 添加：1 消费：2 回滚：3
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
}

var poolPointOperateDto = sync.Pool{
	New: func() any {
		return new(PointOperateDto)
	},
}

// GetPointOperateDto() 从对象池中获取PointOperateDto
func GetPointOperateDto() *PointOperateDto {
	return poolPointOperateDto.Get().(*PointOperateDto)
}

// ReleasePointOperateDto 释放PointOperateDto
func ReleasePointOperateDto(v *PointOperateDto) {
	v.TotalPoint = ""
	v.SceneCode = ""
	v.Remark = ""
	v.LevelPoint = ""
	v.SerialNo = ""
	v.ExtraInfo = ""
	v.UserSite = ""
	v.OpenUserId = ""
	v.ExpireTime = 0
	v.OperateType = 0
	poolPointOperateDto.Put(v)
}
