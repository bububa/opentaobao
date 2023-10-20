package flight

import (
	"sync"
)

// IssueList 结构体
type IssueList struct {
	// 票号
	Tickets []string `json:"tickets,omitempty" xml:"tickets>string,omitempty"`
	// 税项对象
	Taxes []Taxes `json:"taxes,omitempty" xml:"taxes>taxes,omitempty"`
	// 航段
	Segments []Segments `json:"segments,omitempty" xml:"segments>segments,omitempty"`
	// 政策信息
	SellPolicyList []SellPolicyDto `json:"sell_policy_list,omitempty" xml:"sell_policy_list>sell_policy_dto,omitempty"`
	// 证件信息
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 乘机人姓名
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// pnr
	Pnr string `json:"pnr,omitempty" xml:"pnr,omitempty"`
	// 联系电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 乘机人生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 大编码
	BigPnr string `json:"big_pnr,omitempty" xml:"big_pnr,omitempty"`
	// 姓氏
	SurName string `json:"sur_name,omitempty" xml:"sur_name,omitempty"`
	// 名字
	GivenName string `json:"given_name,omitempty" xml:"given_name,omitempty"`
	// 乘机人证件有效期
	CertPeriod string `json:"cert_period,omitempty" xml:"cert_period,omitempty"`
	// 乘机人国籍
	Nationality string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// 乘机人证件颁发国家
	CertIssueCountry string `json:"cert_issue_country,omitempty" xml:"cert_issue_country,omitempty"`
	// 0 身份证 1 护照 2 学生证 3 军官证 4 回乡证 5 台胞证 6 港澳通行证 7 国际海员 8 外国人永久居留证 9 其他证件 10 警官证 11 士兵证件 12 台湾通行证 13 入台证 14 户口簿 15 出生证 16 驾驶证 17 港澳居民居住证 18 台湾居民居住证
	CertType int64 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 乘客类型: 1:成人, 2:儿童, 3:婴儿, 4:留学生
	PassengerType int64 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 票面价
	TicketPrice int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 优惠价格
	Promotion int64 `json:"promotion,omitempty" xml:"promotion,omitempty"`
	// 乘机人性别:1表示男性，2表示女性
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
}

var poolIssueList = sync.Pool{
	New: func() any {
		return new(IssueList)
	},
}

// GetIssueList() 从对象池中获取IssueList
func GetIssueList() *IssueList {
	return poolIssueList.Get().(*IssueList)
}

// ReleaseIssueList 释放IssueList
func ReleaseIssueList(v *IssueList) {
	v.Tickets = v.Tickets[:0]
	v.Taxes = v.Taxes[:0]
	v.Segments = v.Segments[:0]
	v.SellPolicyList = v.SellPolicyList[:0]
	v.CertNo = ""
	v.PassengerName = ""
	v.Pnr = ""
	v.Mobile = ""
	v.Birthday = ""
	v.BigPnr = ""
	v.SurName = ""
	v.GivenName = ""
	v.CertPeriod = ""
	v.Nationality = ""
	v.CertIssueCountry = ""
	v.CertType = 0
	v.PassengerType = 0
	v.TicketPrice = 0
	v.Promotion = 0
	v.Gender = 0
	poolIssueList.Put(v)
}
