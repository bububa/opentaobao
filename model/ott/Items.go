package ott

import (
	"sync"
)

// Items 结构体
type Items struct {
	// 便签按钮
	NoteBtns []Notebtns `json:"note_btns,omitempty" xml:"note_btns>notebtns,omitempty"`
	// 爆米花图
	PicPops []string `json:"pic_pops,omitempty" xml:"pic_pops>string,omitempty"`
	// 坑位列表
	ItemList []Itemlist `json:"item_list,omitempty" xml:"item_list>itemlist,omitempty"`
	// 类目入口
	EntryList []Entrylist `json:"entry_list,omitempty" xml:"entry_list>entrylist,omitempty"`
	// 副标题
	SubTitles []string `json:"sub_titles,omitempty" xml:"sub_titles>string,omitempty"`
	// 待推荐主题
	RuleIds []string `json:"rule_ids,omitempty" xml:"rule_ids>string,omitempty"`
	// 动画类型
	AnimeType string `json:"anime_type,omitempty" xml:"anime_type,omitempty"`
	// 主标题
	MainTitle string `json:"main_title,omitempty" xml:"main_title,omitempty"`
	// 前景图标
	PicUrl1 string `json:"pic_url1,omitempty" xml:"pic_url1,omitempty"`
	// 动画图标
	PicUrl2 string `json:"pic_url2,omitempty" xml:"pic_url2,omitempty"`
	// 背景图标
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 摘要看点
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 坑位类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 坑位名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 坑位校验码
	VerifyCode string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
	// 展示方式
	DisplayType string `json:"display_type,omitempty" xml:"display_type,omitempty"`
	// 标题图
	PicTitle string `json:"pic_title,omitempty" xml:"pic_title,omitempty"`
	// 行为扩展
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 跳转行为
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 显示方式
	ShowType string `json:"show_type,omitempty" xml:"show_type,omitempty"`
	// 背景图标
	PicMap string `json:"pic_map,omitempty" xml:"pic_map,omitempty"`
	// 坑位号
	Position int64 `json:"position,omitempty" xml:"position,omitempty"`
	// 角标图
	PicCorner *PicCornerDo `json:"pic_corner,omitempty" xml:"pic_corner,omitempty"`
	// 内容坑位ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 轮询时间
	CycleTime int64 `json:"cycle_time,omitempty" xml:"cycle_time,omitempty"`
	// 桌面坑位ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 功能入口
	Entry *ItemEntryDo `json:"entry,omitempty" xml:"entry,omitempty"`
	// 排行榜
	Chart *MetaChartDo `json:"chart,omitempty" xml:"chart,omitempty"`
	// 轮播频道ID
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 入口风格
	EntryStyle int64 `json:"entry_style,omitempty" xml:"entry_style,omitempty"`
	// 推荐主题
	RuleId int64 `json:"rule_id,omitempty" xml:"rule_id,omitempty"`
	// 推荐场景
	RecAppId int64 `json:"rec_app_id,omitempty" xml:"rec_app_id,omitempty"`
	// 推荐类型,0:运营自主推荐,1:纯个性化推荐,2:主题个性化推荐
	RecType int64 `json:"rec_type,omitempty" xml:"rec_type,omitempty"`
}

var poolItems = sync.Pool{
	New: func() any {
		return new(Items)
	},
}

// GetItems() 从对象池中获取Items
func GetItems() *Items {
	return poolItems.Get().(*Items)
}

// ReleaseItems 释放Items
func ReleaseItems(v *Items) {
	v.NoteBtns = v.NoteBtns[:0]
	v.PicPops = v.PicPops[:0]
	v.ItemList = v.ItemList[:0]
	v.EntryList = v.EntryList[:0]
	v.SubTitles = v.SubTitles[:0]
	v.RuleIds = v.RuleIds[:0]
	v.AnimeType = ""
	v.MainTitle = ""
	v.PicUrl1 = ""
	v.PicUrl2 = ""
	v.PicUrl = ""
	v.Summary = ""
	v.Type = ""
	v.Name = ""
	v.VerifyCode = ""
	v.DisplayType = ""
	v.PicTitle = ""
	v.Extra = ""
	v.Action = ""
	v.ShowType = ""
	v.PicMap = ""
	v.Position = 0
	v.PicCorner = nil
	v.Id = 0
	v.CycleTime = 0
	v.ItemId = 0
	v.Entry = nil
	v.Chart = nil
	v.ChannelId = 0
	v.EntryStyle = 0
	v.RuleId = 0
	v.RecAppId = 0
	v.RecType = 0
	poolItems.Put(v)
}
