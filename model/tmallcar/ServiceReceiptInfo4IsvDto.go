package tmallcar

import (
	"sync"
)

// ServiceReceiptInfo4IsvDto 结构体
type ServiceReceiptInfo4IsvDto struct {
	// 具体地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 品牌
	Brand string `json:"brand,omitempty" xml:"brand,omitempty"`
	// 选择问题列表，多个问题以英文逗号分隔
	ChooseProblems string `json:"choose_problems,omitempty" xml:"choose_problems,omitempty"`
	// 选定的门店ID
	ChosenOuterStoreId string `json:"chosen_outer_store_id,omitempty" xml:"chosen_outer_store_id,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	County string `json:"county,omitempty" xml:"county,omitempty"`
	// 工单附加信息
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 问题描述图片列表，多个图片以英文逗号分隔
	ProblemDescPics string `json:"problem_desc_pics,omitempty" xml:"problem_desc_pics,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 用户备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// DOOR2DOOR:上门服务，INSTORE:到店服务
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 工单状态： SUBMIT:预约成功，待服务 CANCEL:已取消 FINISH:服务已完成
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 店铺名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 古荡街道
	Town string `json:"town,omitempty" xml:"town,omitempty"`
	// GUARANTEE:维修，INSTALL:安装
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 消费者名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 工单号
	ReceiptId int64 `json:"receipt_id,omitempty" xml:"receipt_id,omitempty"`
}

var poolServiceReceiptInfo4IsvDto = sync.Pool{
	New: func() any {
		return new(ServiceReceiptInfo4IsvDto)
	},
}

// GetServiceReceiptInfo4IsvDto() 从对象池中获取ServiceReceiptInfo4IsvDto
func GetServiceReceiptInfo4IsvDto() *ServiceReceiptInfo4IsvDto {
	return poolServiceReceiptInfo4IsvDto.Get().(*ServiceReceiptInfo4IsvDto)
}

// ReleaseServiceReceiptInfo4IsvDto 释放ServiceReceiptInfo4IsvDto
func ReleaseServiceReceiptInfo4IsvDto(v *ServiceReceiptInfo4IsvDto) {
	v.Address = ""
	v.Brand = ""
	v.ChooseProblems = ""
	v.ChosenOuterStoreId = ""
	v.City = ""
	v.County = ""
	v.Extension = ""
	v.Mobile = ""
	v.ProblemDescPics = ""
	v.Province = ""
	v.Remark = ""
	v.ServiceType = ""
	v.Status = ""
	v.StoreName = ""
	v.Town = ""
	v.Type = ""
	v.UserName = ""
	v.ReceiptId = 0
	poolServiceReceiptInfo4IsvDto.Put(v)
}
