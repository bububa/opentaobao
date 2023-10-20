package wms

import (
	"sync"
)

// Receiverinfowlbwmsstockinordernotifywl 结构体
type Receiverinfowlbwmsstockinordernotifywl struct {
	// 收件人手机
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 收件人手机
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收件人名称
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收件方地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收件人区县
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
	// 收件人城市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 收件人省份
	ReceiverProvince string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
	// 收件人邮编
	ReceiverZipCode string `json:"receiver_zip_code,omitempty" xml:"receiver_zip_code,omitempty"`
	// 收件人镇
	ReceiverTown string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
}

var poolReceiverinfowlbwmsstockinordernotifywl = sync.Pool{
	New: func() any {
		return new(Receiverinfowlbwmsstockinordernotifywl)
	},
}

// GetReceiverinfowlbwmsstockinordernotifywl() 从对象池中获取Receiverinfowlbwmsstockinordernotifywl
func GetReceiverinfowlbwmsstockinordernotifywl() *Receiverinfowlbwmsstockinordernotifywl {
	return poolReceiverinfowlbwmsstockinordernotifywl.Get().(*Receiverinfowlbwmsstockinordernotifywl)
}

// ReleaseReceiverinfowlbwmsstockinordernotifywl 释放Receiverinfowlbwmsstockinordernotifywl
func ReleaseReceiverinfowlbwmsstockinordernotifywl(v *Receiverinfowlbwmsstockinordernotifywl) {
	v.ReceiverPhone = ""
	v.ReceiverMobile = ""
	v.ReceiverName = ""
	v.ReceiverAddress = ""
	v.ReceiverArea = ""
	v.ReceiverCity = ""
	v.ReceiverProvince = ""
	v.ReceiverZipCode = ""
	v.ReceiverTown = ""
	poolReceiverinfowlbwmsstockinordernotifywl.Put(v)
}
