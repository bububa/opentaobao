package btrip

import (
	"github.com/bububa/opentaobao/model"
)

// TravelerInfo 结构体
type TravelerInfo struct {
	// 证件号
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 证件类型
	CertType string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 用户名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户编号
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 出发机场三字码
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 到达机场三字码
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 出发机场三字码
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 到达机场三字码
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 乘客类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 生日（非身份证必选）
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 乘机人手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 证件有效期
	CertValidDate string `json:"cert_valid_date,omitempty" xml:"cert_valid_date,omitempty"`
	// 证件签发国
	CertIssueCountry string `json:"cert_issue_country,omitempty" xml:"cert_issue_country,omitempty"`
	// 国籍
	Nationality string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// 国籍二字码
	NationalityCode string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// 证件签发国
	CertNation string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// 性别，0是男，1是女
	Sex *model.File `json:"sex,omitempty" xml:"sex,omitempty"`
}
