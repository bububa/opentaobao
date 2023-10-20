package ascpchannel

import (
	"sync"
)

// AlibabaAscpChannelSalesOrderCreateData 结构体
type AlibabaAscpChannelSalesOrderCreateData struct {
	// 子单列表
	SubOrderList []Suborders `json:"sub_order_list,omitempty" xml:"sub_order_list>suborders,omitempty"`
	// 渠道订单号
	SaleOrderNo string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
}

var poolAlibabaAscpChannelSalesOrderCreateData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelSalesOrderCreateData)
	},
}

// GetAlibabaAscpChannelSalesOrderCreateData() 从对象池中获取AlibabaAscpChannelSalesOrderCreateData
func GetAlibabaAscpChannelSalesOrderCreateData() *AlibabaAscpChannelSalesOrderCreateData {
	return poolAlibabaAscpChannelSalesOrderCreateData.Get().(*AlibabaAscpChannelSalesOrderCreateData)
}

// ReleaseAlibabaAscpChannelSalesOrderCreateData 释放AlibabaAscpChannelSalesOrderCreateData
func ReleaseAlibabaAscpChannelSalesOrderCreateData(v *AlibabaAscpChannelSalesOrderCreateData) {
	v.SubOrderList = v.SubOrderList[:0]
	v.SaleOrderNo = ""
	poolAlibabaAscpChannelSalesOrderCreateData.Put(v)
}
