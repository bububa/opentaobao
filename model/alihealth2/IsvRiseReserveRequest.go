package alihealth2

import (
	"sync"
)

// IsvRiseReserveRequest 结构体
type IsvRiseReserveRequest struct {
	// 科室名称
	DepartmentName string `json:"department_name,omitempty" xml:"department_name,omitempty"`
	// 支付宝订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// ISV门店名称
	SpStoreName string `json:"sp_store_name,omitempty" xml:"sp_store_name,omitempty"`
	// 证件号码
	IdNumber string `json:"id_number,omitempty" xml:"id_number,omitempty"`
	// userId
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// ISV项目名称，多个项目英文逗号分割
	SpItemName string `json:"sp_item_name,omitempty" xml:"sp_item_name,omitempty"`
	// ISV项目ID,多个项目英文逗号分割
	SpItemId string `json:"sp_item_id,omitempty" xml:"sp_item_id,omitempty"`
	// 医生名称
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 用户号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 用户姓名
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// ISV门店ID
	SpStoreId string `json:"sp_store_id,omitempty" xml:"sp_store_id,omitempty"`
	// 预约时间
	ReserveTime string `json:"reserve_time,omitempty" xml:"reserve_time,omitempty"`
	// ISV复诊/升单ID
	SpReserveId string `json:"sp_reserve_id,omitempty" xml:"sp_reserve_id,omitempty"`
	// 证件类型 0身份证 1护照
	IdType int64 `json:"id_type,omitempty" xml:"id_type,omitempty"`
	// 支付宝初诊ID
	FirstReserveId int64 `json:"first_reserve_id,omitempty" xml:"first_reserve_id,omitempty"`
	// 性别 0男1女
	Sex int64 `json:"sex,omitempty" xml:"sex,omitempty"`
	// 0 新增 1 修改
	OpType int64 `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// 是否结婚 0未婚1结婚
	IsMarried int64 `json:"is_married,omitempty" xml:"is_married,omitempty"`
	// 支付宝标品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 修改状态 2 待就诊 4 已就诊
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolIsvRiseReserveRequest = sync.Pool{
	New: func() any {
		return new(IsvRiseReserveRequest)
	},
}

// GetIsvRiseReserveRequest() 从对象池中获取IsvRiseReserveRequest
func GetIsvRiseReserveRequest() *IsvRiseReserveRequest {
	return poolIsvRiseReserveRequest.Get().(*IsvRiseReserveRequest)
}

// ReleaseIsvRiseReserveRequest 释放IsvRiseReserveRequest
func ReleaseIsvRiseReserveRequest(v *IsvRiseReserveRequest) {
	v.DepartmentName = ""
	v.OrderId = ""
	v.SpStoreName = ""
	v.IdNumber = ""
	v.UserId = ""
	v.SpItemName = ""
	v.SpItemId = ""
	v.DoctorName = ""
	v.Phone = ""
	v.UserNick = ""
	v.SpStoreId = ""
	v.ReserveTime = ""
	v.SpReserveId = ""
	v.IdType = 0
	v.FirstReserveId = 0
	v.Sex = 0
	v.OpType = 0
	v.IsMarried = 0
	v.ItemId = 0
	v.Status = 0
	poolIsvRiseReserveRequest.Put(v)
}
