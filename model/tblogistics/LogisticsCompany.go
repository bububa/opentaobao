package tblogistics

import (
	"sync"
)

// LogisticsCompany 结构体
type LogisticsCompany struct {
	// 物流公司代码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 物流公司简称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 运单号验证正则表达式
	RegMailNo string `json:"reg_mail_no,omitempty" xml:"reg_mail_no,omitempty"`
	// 物流公司标识
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolLogisticsCompany = sync.Pool{
	New: func() any {
		return new(LogisticsCompany)
	},
}

// GetLogisticsCompany() 从对象池中获取LogisticsCompany
func GetLogisticsCompany() *LogisticsCompany {
	return poolLogisticsCompany.Get().(*LogisticsCompany)
}

// ReleaseLogisticsCompany 释放LogisticsCompany
func ReleaseLogisticsCompany(v *LogisticsCompany) {
	v.Code = ""
	v.Name = ""
	v.RegMailNo = ""
	v.Id = 0
	poolLogisticsCompany.Put(v)
}
