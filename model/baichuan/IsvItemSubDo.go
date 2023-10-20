package baichuan

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// IsvItemSubDo 结构体
type IsvItemSubDo struct {
	// 添加时间
	AddTime string `json:"add_time,omitempty" xml:"add_time,omitempty"`
	// 商品状态
	ItemStatus *model.File `json:"item_status,omitempty" xml:"item_status,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 数据库索引id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolIsvItemSubDo = sync.Pool{
	New: func() any {
		return new(IsvItemSubDo)
	},
}

// GetIsvItemSubDo() 从对象池中获取IsvItemSubDo
func GetIsvItemSubDo() *IsvItemSubDo {
	return poolIsvItemSubDo.Get().(*IsvItemSubDo)
}

// ReleaseIsvItemSubDo 释放IsvItemSubDo
func ReleaseIsvItemSubDo(v *IsvItemSubDo) {
	v.AddTime = ""
	v.ItemStatus = nil
	v.ItemId = 0
	v.Id = 0
	poolIsvItemSubDo.Put(v)
}
