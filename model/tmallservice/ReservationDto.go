package tmallservice

import (
	"sync"
)

// ReservationDto 结构体
type ReservationDto struct {
	// 预约时间
	ReserveTimeEnd string `json:"reserve_time_end,omitempty" xml:"reserve_time_end,omitempty"`
	// 预约时间
	ReserveTimeStart string `json:"reserve_time_start,omitempty" xml:"reserve_time_start,omitempty"`
	// 预约来源
	ReserveRequestSourceType string `json:"reserve_request_source_type,omitempty" xml:"reserve_request_source_type,omitempty"`
	// 内部订单号
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 天猫订单号列表
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
	// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 预约日期,2015-12-15
	ResvDate string `json:"resv_date,omitempty" xml:"resv_date,omitempty"`
	// 预约工人手机号码
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 工人名称
	WorkerName string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
	// 下次预约时间
	NextResvTime string `json:"next_resv_time,omitempty" xml:"next_resv_time,omitempty"`
	// 身份证信息
	WorkerIdNum string `json:"worker_id_num,omitempty" xml:"worker_id_num,omitempty"`
	// 预约时间段，会取代resv_time的上午/下午/晚上
	ResvAccurateTime string `json:"resv_accurate_time,omitempty" xml:"resv_accurate_time,omitempty"`
	// 预约时间,0:上午,1:下午,2:晚上
	ResvTime int64 `json:"resv_time,omitempty" xml:"resv_time,omitempty"`
	// 1：电话占线/无人接听/电话关机 	 2：未收到货 	 3：用户暂不需要安装 	 4：取消安装 	 5：电话号码错误
	FailCode int64 `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 1成功 0失败
	Success int64 `json:"success,omitempty" xml:"success,omitempty"`
	// 跳过订单类型检查的原因ID, 具体原因对应ID值咨询运营，一般不填写
	StopOrderTypeCheckReason int64 `json:"stop_order_type_check_reason,omitempty" xml:"stop_order_type_check_reason,omitempty"`
}

var poolReservationDto = sync.Pool{
	New: func() any {
		return new(ReservationDto)
	},
}

// GetReservationDto() 从对象池中获取ReservationDto
func GetReservationDto() *ReservationDto {
	return poolReservationDto.Get().(*ReservationDto)
}

// ReleaseReservationDto 释放ReservationDto
func ReleaseReservationDto(v *ReservationDto) {
	v.ReserveTimeEnd = ""
	v.ReserveTimeStart = ""
	v.ReserveRequestSourceType = ""
	v.OuterId = ""
	v.OrderIds = ""
	v.ServiceType = ""
	v.ResvDate = ""
	v.WorkerMobile = ""
	v.WorkerName = ""
	v.NextResvTime = ""
	v.WorkerIdNum = ""
	v.ResvAccurateTime = ""
	v.ResvTime = 0
	v.FailCode = 0
	v.Success = 0
	v.StopOrderTypeCheckReason = 0
	poolReservationDto.Put(v)
}
