package traveltrade

import (
	"sync"
)

// TravellerInfo 结构体
type TravellerInfo struct {
	// 出生日期，格式yyyy-mm-dd
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 证件号码
	CredentialNo string `json:"credential_no,omitempty" xml:"credential_no,omitempty"`
	// 联系电子邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 预留，KV对形式，商家自定义的个性化出行人信息。目前支持的Key列表如下：overseaPhoneNumber（国外手机号），emergencyPhoneNumber（紧急联系方式），memo（备注）
	ExtendAttributes string `json:"extend_attributes,omitempty" xml:"extend_attributes,omitempty"`
	// 证件签发国
	IssueCountry string `json:"issue_country,omitempty" xml:"issue_country,omitempty"`
	// 证件签发地
	IssuePlace string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 姓名（拼音）
	NamePinyin string `json:"name_pinyin,omitempty" xml:"name_pinyin,omitempty"`
	// 国籍
	Nationality string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 具体收件地址（对于需要物流的实物商品）
	PostAddress string `json:"post_address,omitempty" xml:"post_address,omitempty"`
	// 收件区域
	PostArea string `json:"post_area,omitempty" xml:"post_area,omitempty"`
	// 收件市（对于需要物流的实物商品）
	PostCity string `json:"post_city,omitempty" xml:"post_city,omitempty"`
	// 收件省（对于需要物流的实物商品）
	PostProvince string `json:"post_province,omitempty" xml:"post_province,omitempty"`
	// 证件有效期
	ValidDate string `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 姓（拼音）
	SurnamePinyin string `json:"surname_pinyin,omitempty" xml:"surname_pinyin,omitempty"`
	// 名（拼音）
	GivenNamePinyin string `json:"given_name_pinyin,omitempty" xml:"given_name_pinyin,omitempty"`
	// 预留，暂时无用
	ExtendAttributesJson string `json:"extend_attributes_json,omitempty" xml:"extend_attributes_json,omitempty"`
	// 证件类型。0:身份证 1:护照 2:学生证3:军官证 4:回乡证 5:台胞证 6:港澳通行证 10:警官证 11:士兵证 12:台湾通行证
	CredentialType int64 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// 性别。0-男，1-女
	Sex int64 `json:"sex,omitempty" xml:"sex,omitempty"`
}

var poolTravellerInfo = sync.Pool{
	New: func() any {
		return new(TravellerInfo)
	},
}

// GetTravellerInfo() 从对象池中获取TravellerInfo
func GetTravellerInfo() *TravellerInfo {
	return poolTravellerInfo.Get().(*TravellerInfo)
}

// ReleaseTravellerInfo 释放TravellerInfo
func ReleaseTravellerInfo(v *TravellerInfo) {
	v.Birthday = ""
	v.CredentialNo = ""
	v.Email = ""
	v.ExtendAttributes = ""
	v.IssueCountry = ""
	v.IssuePlace = ""
	v.Name = ""
	v.NamePinyin = ""
	v.Nationality = ""
	v.Phone = ""
	v.PostAddress = ""
	v.PostArea = ""
	v.PostCity = ""
	v.PostProvince = ""
	v.ValidDate = ""
	v.SurnamePinyin = ""
	v.GivenNamePinyin = ""
	v.ExtendAttributesJson = ""
	v.CredentialType = 0
	v.Sex = 0
	poolTravellerInfo.Put(v)
}
