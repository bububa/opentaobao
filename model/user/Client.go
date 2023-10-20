package user

import (
	"sync"
)

// Client 结构体
type Client struct {
	// 联系人列表
	ContactPersonList []ContactPerson `json:"contact_person_list,omitempty" xml:"contact_person_list>contact_person,omitempty"`
	// 客户登陆账号
	MerchantLoginId string `json:"merchant_login_id,omitempty" xml:"merchant_login_id,omitempty"`
}

var poolClient = sync.Pool{
	New: func() any {
		return new(Client)
	},
}

// GetClient() 从对象池中获取Client
func GetClient() *Client {
	return poolClient.Get().(*Client)
}

// ReleaseClient 释放Client
func ReleaseClient(v *Client) {
	v.ContactPersonList = v.ContactPersonList[:0]
	v.MerchantLoginId = ""
	poolClient.Put(v)
}
