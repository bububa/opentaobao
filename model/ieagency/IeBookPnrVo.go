package ieagency

import (
	"sync"
)

// IeBookPnrVo 结构体
type IeBookPnrVo struct {
	// pnr值
	PnrNo string `json:"pnr_no,omitempty" xml:"pnr_no,omitempty"`
	// pnr类型
	PnrType string `json:"pnr_type,omitempty" xml:"pnr_type,omitempty"`
	// pnr id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolIeBookPnrVo = sync.Pool{
	New: func() any {
		return new(IeBookPnrVo)
	},
}

// GetIeBookPnrVo() 从对象池中获取IeBookPnrVo
func GetIeBookPnrVo() *IeBookPnrVo {
	return poolIeBookPnrVo.Get().(*IeBookPnrVo)
}

// ReleaseIeBookPnrVo 释放IeBookPnrVo
func ReleaseIeBookPnrVo(v *IeBookPnrVo) {
	v.PnrNo = ""
	v.PnrType = ""
	v.Id = 0
	poolIeBookPnrVo.Put(v)
}
