package alimember

import (
	"sync"
)

// SyncMemberPointChangeDto 结构体
type SyncMemberPointChangeDto struct {
	// 主账号使用main，子账号使用sub
	UidType string `json:"uid_type,omitempty" xml:"uid_type,omitempty"`
	// 淘系交易的订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// isv存储的会员id，唯一标识一个会员
	OuterMemberId string `json:"outer_member_id,omitempty" xml:"outer_member_id,omitempty"`
	// 成长值变化量
	RawQuantity string `json:"raw_quantity,omitempty" xml:"raw_quantity,omitempty"`
	// 变更同步的序列号，幂等使用
	SerialNo string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`
	// 用户的ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 变更前的成长值
	OldTotalPoint string `json:"old_total_point,omitempty" xml:"old_total_point,omitempty"`
	// 变更后的成长值
	TotalPoint string `json:"total_point,omitempty" xml:"total_point,omitempty"`
	// 变更类型 1增加 2减少
	OperateType int64 `json:"operate_type,omitempty" xml:"operate_type,omitempty"`
	// 变更渠道 1淘宝 100其他
	Channel int64 `json:"channel,omitempty" xml:"channel,omitempty"`
	// 变更场景 0交易 1互动活动 100其他
	BizScene int64 `json:"biz_scene,omitempty" xml:"biz_scene,omitempty"`
	// 积分类型 1消费积分 2成长值(等级积分)
	PointType int64 `json:"point_type,omitempty" xml:"point_type,omitempty"`
	// 成长值变更发生时间，毫秒级时间戳，一般早于同步时间
	ChangeTime int64 `json:"change_time,omitempty" xml:"change_time,omitempty"`
}

var poolSyncMemberPointChangeDto = sync.Pool{
	New: func() any {
		return new(SyncMemberPointChangeDto)
	},
}

// GetSyncMemberPointChangeDto() 从对象池中获取SyncMemberPointChangeDto
func GetSyncMemberPointChangeDto() *SyncMemberPointChangeDto {
	return poolSyncMemberPointChangeDto.Get().(*SyncMemberPointChangeDto)
}

// ReleaseSyncMemberPointChangeDto 释放SyncMemberPointChangeDto
func ReleaseSyncMemberPointChangeDto(v *SyncMemberPointChangeDto) {
	v.UidType = ""
	v.OrderId = ""
	v.OuterMemberId = ""
	v.RawQuantity = ""
	v.SerialNo = ""
	v.Ouid = ""
	v.OldTotalPoint = ""
	v.TotalPoint = ""
	v.OperateType = 0
	v.Channel = 0
	v.BizScene = 0
	v.PointType = 0
	v.ChangeTime = 0
	poolSyncMemberPointChangeDto.Put(v)
}
