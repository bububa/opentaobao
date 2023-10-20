package ieagency

import (
	"sync"
)

// IeBookTicketVo 结构体
type IeBookTicketVo struct {
	// 航班号id列表
	FlightIds []int64 `json:"flight_ids,omitempty" xml:"flight_ids>int64,omitempty"`
	// 票号
	TicketNos []string `json:"ticket_nos,omitempty" xml:"ticket_nos>string,omitempty"`
	// 票号验真备注
	AuthenticMemo string `json:"authentic_memo,omitempty" xml:"authentic_memo,omitempty"`
	// 票号验真状态(Init:初始状态,Failure:验真失败,SUCCESS:验真通过)
	AuthenticStatus string `json:"authentic_status,omitempty" xml:"authentic_status,omitempty"`
	// 出票日期
	IssueTicketTime string `json:"issue_ticket_time,omitempty" xml:"issue_ticket_time,omitempty"`
	// 乘机人证件号
	PassengerCertNo string `json:"passenger_cert_no,omitempty" xml:"passenger_cert_no,omitempty"`
	// 乘机人证件类型(Passport:护照,Hongkong_Macao:港澳通行证,Taiwan_MTP:台胞证,Home_Return_Permit:回乡证,Taiwan_Pass:台湾通行证,Entry_Taiwan_Permit:入台证);
	PassengerCertType string `json:"passenger_cert_type,omitempty" xml:"passenger_cert_type,omitempty"`
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 乘机人类型(Adult:普通成人,Child:儿童,StudentAbroad:留学生,Infant:婴儿)
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// pnr主键
	PnrId int64 `json:"pnr_id,omitempty" xml:"pnr_id,omitempty"`
	// 票价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 票税
	TicketTax int64 `json:"ticket_tax,omitempty" xml:"ticket_tax,omitempty"`
	// 是否联程票
	ConjTicketFlag bool `json:"conj_ticket_flag,omitempty" xml:"conj_ticket_flag,omitempty"`
}

var poolIeBookTicketVo = sync.Pool{
	New: func() any {
		return new(IeBookTicketVo)
	},
}

// GetIeBookTicketVo() 从对象池中获取IeBookTicketVo
func GetIeBookTicketVo() *IeBookTicketVo {
	return poolIeBookTicketVo.Get().(*IeBookTicketVo)
}

// ReleaseIeBookTicketVo 释放IeBookTicketVo
func ReleaseIeBookTicketVo(v *IeBookTicketVo) {
	v.FlightIds = v.FlightIds[:0]
	v.TicketNos = v.TicketNos[:0]
	v.AuthenticMemo = ""
	v.AuthenticStatus = ""
	v.IssueTicketTime = ""
	v.PassengerCertNo = ""
	v.PassengerCertType = ""
	v.PassengerName = ""
	v.PassengerType = ""
	v.Id = 0
	v.PnrId = 0
	v.TicketPrice = 0
	v.TicketTax = 0
	v.ConjTicketFlag = false
	poolIeBookTicketVo.Put(v)
}
