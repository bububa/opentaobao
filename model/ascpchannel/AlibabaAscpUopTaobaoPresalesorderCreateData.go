package ascpchannel

import (
	"sync"
)

// AlibabaAscpUopTaobaoPresalesorderCreateData 结构体
type AlibabaAscpUopTaobaoPresalesorderCreateData struct {
	// 物流信息
	LogisticsAcceptInfoResList []Logisticsacceptinforeslist `json:"logistics_accept_info_res_list,omitempty" xml:"logistics_accept_info_res_list>logisticsacceptinforeslist,omitempty"`
}

var poolAlibabaAscpUopTaobaoPresalesorderCreateData = sync.Pool{
	New: func() any {
		return new(AlibabaAscpUopTaobaoPresalesorderCreateData)
	},
}

// GetAlibabaAscpUopTaobaoPresalesorderCreateData() 从对象池中获取AlibabaAscpUopTaobaoPresalesorderCreateData
func GetAlibabaAscpUopTaobaoPresalesorderCreateData() *AlibabaAscpUopTaobaoPresalesorderCreateData {
	return poolAlibabaAscpUopTaobaoPresalesorderCreateData.Get().(*AlibabaAscpUopTaobaoPresalesorderCreateData)
}

// ReleaseAlibabaAscpUopTaobaoPresalesorderCreateData 释放AlibabaAscpUopTaobaoPresalesorderCreateData
func ReleaseAlibabaAscpUopTaobaoPresalesorderCreateData(v *AlibabaAscpUopTaobaoPresalesorderCreateData) {
	v.LogisticsAcceptInfoResList = v.LogisticsAcceptInfoResList[:0]
	poolAlibabaAscpUopTaobaoPresalesorderCreateData.Put(v)
}
