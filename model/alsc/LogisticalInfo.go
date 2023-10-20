package alsc

import (
	"sync"
)

// LogisticalInfo 结构体
type LogisticalInfo struct {
	// 送达时间
	ArriveTime string `json:"arrive_time,omitempty" xml:"arrive_time,omitempty"`
	// 物流状态：WAIT_DELIVERY(&#34;WAIT_DELIVERY&#34;, &#34;待发货&#34;),  WAIT_RECEIVE(&#34;WAIT_RECEIVE&#34;, &#34;待收货&#34;),  SUCCESS(&#34;SUCCESS&#34;, &#34;确认收货&#34;),  REFUND(&#34;REFUND&#34;, &#34;退货&#34;);
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
	// 收款地址
	ReceiveAddress string `json:"receive_address,omitempty" xml:"receive_address,omitempty"`
	// 联系人手机
	ReceivePhone string `json:"receive_phone,omitempty" xml:"receive_phone,omitempty"`
}

var poolLogisticalInfo = sync.Pool{
	New: func() any {
		return new(LogisticalInfo)
	},
}

// GetLogisticalInfo() 从对象池中获取LogisticalInfo
func GetLogisticalInfo() *LogisticalInfo {
	return poolLogisticalInfo.Get().(*LogisticalInfo)
}

// ReleaseLogisticalInfo 释放LogisticalInfo
func ReleaseLogisticalInfo(v *LogisticalInfo) {
	v.ArriveTime = ""
	v.LogisticsStatus = ""
	v.ReceiveAddress = ""
	v.ReceivePhone = ""
	poolLogisticalInfo.Put(v)
}
