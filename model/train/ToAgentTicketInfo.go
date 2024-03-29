package train

import (
	"sync"
)

// ToAgentTicketInfo 结构体
type ToAgentTicketInfo struct {
	// 淘宝火车票子订单id.
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 出发站
	FromStation string `json:"from_station,omitempty" xml:"from_station,omitempty"`
	// 出发时间
	FromTime string `json:"from_time,omitempty" xml:"from_time,omitempty"`
	// 到达站
	ToStation string `json:"to_station,omitempty" xml:"to_station,omitempty"`
	// 车次
	TrainNum string `json:"train_num,omitempty" xml:"train_num,omitempty"`
	// 乘客姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 证件编号
	CertificateNum string `json:"certificate_num,omitempty" xml:"certificate_num,omitempty"`
	// 证件类型，0:身份证 1:护照 4:港澳通行证 5:台湾通行证
	CertificateType string `json:"certificate_type,omitempty" xml:"certificate_type,omitempty"`
	// 乘客生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 到站时间
	ToTime string `json:"to_time,omitempty" xml:"to_time,omitempty"`
	// 证件有效期
	ValidUntil string `json:"valid_until,omitempty" xml:"valid_until,omitempty"`
	// 国家code
	NationalityCode string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// 国家名称
	Nationality string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// 性别 1：男 0：女
	Gender string `json:"gender,omitempty" xml:"gender,omitempty"`
	// 联系人电话
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 学生信息
	StudentInfo *StudentInfo `json:"student_info,omitempty" xml:"student_info,omitempty"`
	// 坐席
	Seat int64 `json:"seat,omitempty" xml:"seat,omitempty"`
	// 保险价格，精确到分，例如10元，输入1000。
	InsurancePrice int64 `json:"insurance_price,omitempty" xml:"insurance_price,omitempty"`
	// 单张票价(不包含保险价格),例如100元,输出为10000,精确到分.
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 1:单程票
	Tag int64 `json:"tag,omitempty" xml:"tag,omitempty"`
	// 保险的单一价格
	InsuranceUnitPrice int64 `json:"insurance_unit_price,omitempty" xml:"insurance_unit_price,omitempty"`
	// 0:成人 1:儿童 2：学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// segmentIndex
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 实际票价(不包含保险价格),例如100元,输出为10000,精确到分.
	RealTicketPrice int64 `json:"real_ticket_price,omitempty" xml:"real_ticket_price,omitempty"`
	// 是否支持无座
	SupportNoSeat bool `json:"support_no_seat,omitempty" xml:"support_no_seat,omitempty"`
}

var poolToAgentTicketInfo = sync.Pool{
	New: func() any {
		return new(ToAgentTicketInfo)
	},
}

// GetToAgentTicketInfo() 从对象池中获取ToAgentTicketInfo
func GetToAgentTicketInfo() *ToAgentTicketInfo {
	return poolToAgentTicketInfo.Get().(*ToAgentTicketInfo)
}

// ReleaseToAgentTicketInfo 释放ToAgentTicketInfo
func ReleaseToAgentTicketInfo(v *ToAgentTicketInfo) {
	v.SubOrderId = ""
	v.FromStation = ""
	v.FromTime = ""
	v.ToStation = ""
	v.TrainNum = ""
	v.PassengerName = ""
	v.CertificateNum = ""
	v.CertificateType = ""
	v.Birthday = ""
	v.ToTime = ""
	v.ValidUntil = ""
	v.NationalityCode = ""
	v.Nationality = ""
	v.Gender = ""
	v.Telephone = ""
	v.StudentInfo = nil
	v.Seat = 0
	v.InsurancePrice = 0
	v.TicketPrice = 0
	v.Tag = 0
	v.InsuranceUnitPrice = 0
	v.PassengerType = 0
	v.SegmentIndex = 0
	v.RealTicketPrice = 0
	v.SupportNoSeat = false
	poolToAgentTicketInfo.Put(v)
}
