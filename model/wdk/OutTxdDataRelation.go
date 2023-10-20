package wdk

import (
	"sync"
)

// OutTxdDataRelation 结构体
type OutTxdDataRelation struct {
	// 子数据类型
	SubBizCode string `json:"sub_biz_code,omitempty" xml:"sub_biz_code,omitempty"`
	// 数据类型
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 外部数据Id
	OutDataId string `json:"out_data_id,omitempty" xml:"out_data_id,omitempty"`
	// 淘鲜达数据Id
	TxdDataId string `json:"txd_data_id,omitempty" xml:"txd_data_id,omitempty"`
}

var poolOutTxdDataRelation = sync.Pool{
	New: func() any {
		return new(OutTxdDataRelation)
	},
}

// GetOutTxdDataRelation() 从对象池中获取OutTxdDataRelation
func GetOutTxdDataRelation() *OutTxdDataRelation {
	return poolOutTxdDataRelation.Get().(*OutTxdDataRelation)
}

// ReleaseOutTxdDataRelation 释放OutTxdDataRelation
func ReleaseOutTxdDataRelation(v *OutTxdDataRelation) {
	v.SubBizCode = ""
	v.BizCode = ""
	v.OutDataId = ""
	v.TxdDataId = ""
	poolOutTxdDataRelation.Put(v)
}
