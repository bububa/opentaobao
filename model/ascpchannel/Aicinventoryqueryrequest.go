package ascpchannel

import (
	"sync"
)

// Aicinventoryqueryrequest 结构体
type Aicinventoryqueryrequest struct {
	// 参数列表
	AicinventoryQueryList []Aicinventoryquerylist `json:"aicinventory_query_list,omitempty" xml:"aicinventory_query_list>aicinventoryquerylist,omitempty"`
}

var poolAicinventoryqueryrequest = sync.Pool{
	New: func() any {
		return new(Aicinventoryqueryrequest)
	},
}

// GetAicinventoryqueryrequest() 从对象池中获取Aicinventoryqueryrequest
func GetAicinventoryqueryrequest() *Aicinventoryqueryrequest {
	return poolAicinventoryqueryrequest.Get().(*Aicinventoryqueryrequest)
}

// ReleaseAicinventoryqueryrequest 释放Aicinventoryqueryrequest
func ReleaseAicinventoryqueryrequest(v *Aicinventoryqueryrequest) {
	v.AicinventoryQueryList = v.AicinventoryQueryList[:0]
	poolAicinventoryqueryrequest.Put(v)
}
