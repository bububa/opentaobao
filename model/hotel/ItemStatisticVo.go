package hotel

import (
	"sync"
)

// ItemStatisticVo 结构体
type ItemStatisticVo struct {
	// tab信息
	RoomTabInfos []TabInfo `json:"room_tab_infos,omitempty" xml:"room_tab_infos>tab_info,omitempty"`
	// 不同分数的数量
	ScoreInfos []ScoreInfo `json:"score_infos,omitempty" xml:"score_infos>score_info,omitempty"`
	// tab信息
	TabInfos []TabInfo `json:"tab_infos,omitempty" xml:"tab_infos>tab_info,omitempty"`
	// 最佳得分项
	BestItem string `json:"best_item,omitempty" xml:"best_item,omitempty"`
	// 推荐率
	RecommendStr string `json:"recommend_str,omitempty" xml:"recommend_str,omitempty"`
	// 评分描述： 非常好
	ScoreDesc string `json:"score_desc,omitempty" xml:"score_desc,omitempty"`
	// 评分详情，json格式
	ScoreDetail string `json:"score_detail,omitempty" xml:"score_detail,omitempty"`
	// 总评分
	TotalScore string `json:"total_score,omitempty" xml:"total_score,omitempty"`
	// 旅游商品信息
	TravelItemInfo string `json:"travel_item_info,omitempty" xml:"travel_item_info,omitempty"`
	// 五分制标记, 1为五分制， 0为十分制
	IsFiveGrade int64 `json:"is_five_grade,omitempty" xml:"is_five_grade,omitempty"`
	// 评论总数
	RateCnt int64 `json:"rate_cnt,omitempty" xml:"rate_cnt,omitempty"`
	// 有图的评论总数
	RatePicCnt int64 `json:"rate_pic_cnt,omitempty" xml:"rate_pic_cnt,omitempty"`
	// 评分星级
	ScoreLevel int64 `json:"score_level,omitempty" xml:"score_level,omitempty"`
	// source来源 0自采 1共享 21agoda 22艺龙 23tripAdvisor
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 热词显示的行数
	TabShowLines int64 `json:"tab_show_lines,omitempty" xml:"tab_show_lines,omitempty"`
	// 旅游商品信息id
	TravelItemId int64 `json:"travel_item_id,omitempty" xml:"travel_item_id,omitempty"`
	// tripAdv评论数
	TripAdvateCnt int64 `json:"trip_advate_cnt,omitempty" xml:"trip_advate_cnt,omitempty"`
}

var poolItemStatisticVo = sync.Pool{
	New: func() any {
		return new(ItemStatisticVo)
	},
}

// GetItemStatisticVo() 从对象池中获取ItemStatisticVo
func GetItemStatisticVo() *ItemStatisticVo {
	return poolItemStatisticVo.Get().(*ItemStatisticVo)
}

// ReleaseItemStatisticVo 释放ItemStatisticVo
func ReleaseItemStatisticVo(v *ItemStatisticVo) {
	v.RoomTabInfos = v.RoomTabInfos[:0]
	v.ScoreInfos = v.ScoreInfos[:0]
	v.TabInfos = v.TabInfos[:0]
	v.BestItem = ""
	v.RecommendStr = ""
	v.ScoreDesc = ""
	v.ScoreDetail = ""
	v.TotalScore = ""
	v.TravelItemInfo = ""
	v.IsFiveGrade = 0
	v.RateCnt = 0
	v.RatePicCnt = 0
	v.ScoreLevel = 0
	v.Source = 0
	v.TabShowLines = 0
	v.TravelItemId = 0
	v.TripAdvateCnt = 0
	poolItemStatisticVo.Put(v)
}
