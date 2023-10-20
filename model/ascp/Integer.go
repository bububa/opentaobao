package ascp

import (
	"sync"
)

// Integer 结构体
type Integer struct {
	// 货品模型
	ItemInfoList []HiErpItemInfo `json:"item_info_list,omitempty" xml:"item_info_list>hi_erp_item_info,omitempty"`
	// 分销商公司/用户名
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 分销商用户id
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}

var poolInteger = sync.Pool{
	New: func() any {
		return new(Integer)
	},
}

// GetInteger() 从对象池中获取Integer
func GetInteger() *Integer {
	return poolInteger.Get().(*Integer)
}

// ReleaseInteger 释放Integer
func ReleaseInteger(v *Integer) {
	v.ItemInfoList = v.ItemInfoList[:0]
	v.CompanyName = ""
	v.MailNo = ""
	v.OwnerId = 0
	poolInteger.Put(v)
}
