package hotel

import (
	"sync"
)

// GetMixRateListParam 结构体
type GetMixRateListParam struct {
	// 酒店评论类型筛选
	TabFilter string `json:"tab_filter,omitempty" xml:"tab_filter,omitempty"`
	// 业务类型
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 用于嵌入页截断显示
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 是否加载顶踩数据，为0不加载，为1加载
	LoadAttitude int64 `json:"load_attitude,omitempty" xml:"load_attitude,omitempty"`
	// 是否加载回复数据，为0不加载，为1加载
	LoadReply int64 `json:"load_reply,omitempty" xml:"load_reply,omitempty"`
	// 页面号，第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页面包含的记录数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolGetMixRateListParam = sync.Pool{
	New: func() any {
		return new(GetMixRateListParam)
	},
}

// GetGetMixRateListParam() 从对象池中获取GetMixRateListParam
func GetGetMixRateListParam() *GetMixRateListParam {
	return poolGetMixRateListParam.Get().(*GetMixRateListParam)
}

// ReleaseGetMixRateListParam 释放GetMixRateListParam
func ReleaseGetMixRateListParam(v *GetMixRateListParam) {
	v.TabFilter = ""
	v.ItemId = 0
	v.Limit = 0
	v.LoadAttitude = 0
	v.LoadReply = 0
	v.PageNo = 0
	v.PageSize = 0
	poolGetMixRateListParam.Put(v)
}
