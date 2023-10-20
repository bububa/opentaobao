package maitix

import (
	"sync"
)

// EticketDto 结构体
type EticketDto struct {
	// 证件号
	CertificateNo string `json:"certificate_no,omitempty" xml:"certificate_no,omitempty"`
	// 证件类型 1:身份证; 2: 护照; 3:港澳居民来往内地通行证; 4:台湾居民来往大陆通行证; 5:士兵/军官证
	CertificateType string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 客户姓名
	CustomerName string `json:"customer_name,omitempty" xml:"customer_name,omitempty"`
	// 入场方式 1 - 电子票 2 - 纸质票 3- 电子票+纸质票
	EntryType string `json:"entry_type,omitempty" xml:"entry_type,omitempty"`
	// 换票码
	ExchangeCode string `json:"exchange_code,omitempty" xml:"exchange_code,omitempty"`
	// 换票方式：1：电子票；2：纸质票
	ExchangeTicketMethod string `json:"exchange_ticket_method,omitempty" xml:"exchange_ticket_method,omitempty"`
	// 楼层名称
	FloorName string `json:"floor_name,omitempty" xml:"floor_name,omitempty"`
	// 楼区排座
	FullSeatName string `json:"full_seat_name,omitempty" xml:"full_seat_name,omitempty"`
	// 演出时间
	PerformTime string `json:"perform_time,omitempty" xml:"perform_time,omitempty"`
	// 联系电话的国家区号
	PhoneCountryCode string `json:"phone_country_code,omitempty" xml:"phone_country_code,omitempty"`
	// 联系电话
	PhoneNumber string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	// 打印状态 ：1、未打印 2、已打印
	PrintStatus string `json:"print_status,omitempty" xml:"print_status,omitempty"`
	// 参数projectName
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 二维码原值
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 二维码类型 1、静态二维码 2、动态二维码3-身份证电子票，4-换票码，5-静态码+换票码，如果是4或5，需要展示给用户exchange_code的原值 6,展示qr_code的值
	QrCodeType string `json:"qr_code_type,omitempty" xml:"qr_code_type,omitempty"`
	// 是否对号入座（入座方式）0=不对号入座 1=对号入座 2=对区入座
	ReserveSeat string `json:"reserve_seat,omitempty" xml:"reserve_seat,omitempty"`
	// 座位：号
	SeatCol string `json:"seat_col,omitempty" xml:"seat_col,omitempty"`
	// 座位：排
	SeatRow string `json:"seat_row,omitempty" xml:"seat_row,omitempty"`
	// 服务电话
	ServicePhone string `json:"service_phone,omitempty" xml:"service_phone,omitempty"`
	// 看台名称
	StandName string `json:"stand_name,omitempty" xml:"stand_name,omitempty"`
	// 看台入场口
	StandPortal string `json:"stand_portal,omitempty" xml:"stand_portal,omitempty"`
	// 票品名称
	TicketItemName string `json:"ticket_item_name,omitempty" xml:"ticket_item_name,omitempty"`
	// 验票状态 1:未验; 2:已验
	ValidateStatus string `json:"validate_status,omitempty" xml:"validate_status,omitempty"`
	// 场馆名称
	VenueName string `json:"venue_name,omitempty" xml:"venue_name,omitempty"`
	// 票池批次号
	BatchCode int64 `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// 套票分组Id
	CombineId int64 `json:"combine_id,omitempty" xml:"combine_id,omitempty"`
	// 楼层Id
	FloorId int64 `json:"floor_id,omitempty" xml:"floor_id,omitempty"`
	// 主分销单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 子分销单Id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 场次ID
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 项目来源: ：1票务云；2直连（第三方）
	ProductSource int64 `json:"product_source,omitempty" xml:"product_source,omitempty"`
	// 项目ID
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 座位id
	SeatId int64 `json:"seat_id,omitempty" xml:"seat_id,omitempty"`
	// 来源系统1 B端，2 C端 默认C端
	SourceSystem int64 `json:"source_system,omitempty" xml:"source_system,omitempty"`
	// 原价格
	SourceTicketItemPrice int64 `json:"source_ticket_item_price,omitempty" xml:"source_ticket_item_price,omitempty"`
	// 项目的代理商家Id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 普通票品价格
	TicketItemPrice int64 `json:"ticket_item_price,omitempty" xml:"ticket_item_price,omitempty"`
	// 票品Id
	TicketItemid int64 `json:"ticket_itemid,omitempty" xml:"ticket_itemid,omitempty"`
	// 场馆id
	VenueId int64 `json:"venue_id,omitempty" xml:"venue_id,omitempty"`
	// 票单号
	VoucherId int64 `json:"voucher_id,omitempty" xml:"voucher_id,omitempty"`
}

var poolEticketDto = sync.Pool{
	New: func() any {
		return new(EticketDto)
	},
}

// GetEticketDto() 从对象池中获取EticketDto
func GetEticketDto() *EticketDto {
	return poolEticketDto.Get().(*EticketDto)
}

// ReleaseEticketDto 释放EticketDto
func ReleaseEticketDto(v *EticketDto) {
	v.CertificateNo = ""
	v.CertificateType = ""
	v.CustomerName = ""
	v.EntryType = ""
	v.ExchangeCode = ""
	v.ExchangeTicketMethod = ""
	v.FloorName = ""
	v.FullSeatName = ""
	v.PerformTime = ""
	v.PhoneCountryCode = ""
	v.PhoneNumber = ""
	v.PrintStatus = ""
	v.ProjectName = ""
	v.QrCode = ""
	v.QrCodeType = ""
	v.ReserveSeat = ""
	v.SeatCol = ""
	v.SeatRow = ""
	v.ServicePhone = ""
	v.StandName = ""
	v.StandPortal = ""
	v.TicketItemName = ""
	v.ValidateStatus = ""
	v.VenueName = ""
	v.BatchCode = 0
	v.CombineId = 0
	v.FloorId = 0
	v.MainOrderId = 0
	v.OrderId = 0
	v.PerformId = 0
	v.ProductSource = 0
	v.ProjectId = 0
	v.SeatId = 0
	v.SourceSystem = 0
	v.SourceTicketItemPrice = 0
	v.SupplierId = 0
	v.TicketItemPrice = 0
	v.TicketItemid = 0
	v.VenueId = 0
	v.VoucherId = 0
	poolEticketDto.Put(v)
}
