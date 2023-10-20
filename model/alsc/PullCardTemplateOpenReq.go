package alsc

import (
	"sync"
)

// PullCardTemplateOpenReq 结构体
type PullCardTemplateOpenReq struct {
	// 品牌id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卡类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 更新时间,数据下行时必须传递
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 上次拉取的业务ID,数据下行时必须传递
	LastMaxId string `json:"last_max_id,omitempty" xml:"last_max_id,omitempty"`
	// 外部品牌id,与brand_id不可同时为空
	OutBrandId string `json:"out_brand_id,omitempty" xml:"out_brand_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 单次返回数量,数据下行时必须传递
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 分页参数,当前页码,分页查询卡模板时必须传递
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页参数,页面大小,分页查询卡模板时必须传递
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否包含逻辑删除,数据下行时必须传递
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 是否需要总数,分页查询卡模板时必须传递
	NeedCount bool `json:"need_count,omitempty" xml:"need_count,omitempty"`
	// 是否查询详情,若需要详情,传true
	DetailRequired bool `json:"detail_required,omitempty" xml:"detail_required,omitempty"`
}

var poolPullCardTemplateOpenReq = sync.Pool{
	New: func() any {
		return new(PullCardTemplateOpenReq)
	},
}

// GetPullCardTemplateOpenReq() 从对象池中获取PullCardTemplateOpenReq
func GetPullCardTemplateOpenReq() *PullCardTemplateOpenReq {
	return poolPullCardTemplateOpenReq.Get().(*PullCardTemplateOpenReq)
}

// ReleasePullCardTemplateOpenReq 释放PullCardTemplateOpenReq
func ReleasePullCardTemplateOpenReq(v *PullCardTemplateOpenReq) {
	v.BrandId = ""
	v.CardType = ""
	v.GmtModified = ""
	v.LastMaxId = ""
	v.OutBrandId = ""
	v.OutShopId = ""
	v.ShopId = ""
	v.Num = 0
	v.PageNo = 0
	v.PageSize = 0
	v.Deleted = false
	v.NeedCount = false
	v.DetailRequired = false
	poolPullCardTemplateOpenReq.Put(v)
}
