package user

// Client 结构体
type Client struct {
	// 联系人列表
	ContactPersonList []ContactPerson `json:"contact_person_list,omitempty" xml:"contact_person_list>contact_person,omitempty"`
	// 客户登陆账号
	MerchantLoginId string `json:"merchant_login_id,omitempty" xml:"merchant_login_id,omitempty"`
}
