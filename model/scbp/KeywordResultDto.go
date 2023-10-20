package scbp

import (
	"sync"
)

// KeywordResultDto 结构体
type KeywordResultDto struct {
	// 关键词所属分组名称列表
	TagList []string `json:"tag_list,omitempty" xml:"tag_list>string,omitempty"`
	// 底价，单位元，保留一位小数, 例如3.5表示3.5元
	BasePrice string `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 购买竞争度[1-6]
	BuyCount string `json:"buy_count,omitempty" xml:"buy_count,omitempty"`
	// 点击量
	ClickCnt string `json:"click_cnt,omitempty" xml:"click_cnt,omitempty"`
	// 平均点击花费，单位元，保留两位小数, 例如3.75表示3.75元
	ClickCostAvg string `json:"click_cost_avg,omitempty" xml:"click_cost_avg,omitempty"`
	// 花费，单位元，保留两位小数, 例如3.75表示3.75元
	Cost string `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击率，百分比，保留两位小数，例如3.75表示3.75%
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 曝光量
	ImpressionCnt string `json:"impression_cnt,omitempty" xml:"impression_cnt,omitempty"`
	// 推广时长，单位小时，保留一位小数，例如13.5表示13.5小时，小于1小时返回&lt;1
	OnlineTime string `json:"online_time,omitempty" xml:"online_time,omitempty"`
	// 出价，单位元，保留一位小数, 例如3.5表示3.5元
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 搜索热度[1-6]
	SearchCount string `json:"search_count,omitempty" xml:"search_count,omitempty"`
	// 关键词推广状态,取值stopped和in_promotion
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 关键词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 和关键词有匹配且处于推广中的产品的个数
	MatchCount int64 `json:"match_count,omitempty" xml:"match_count,omitempty"`
	// 推广评分星级[0-5]
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
	// 关键词id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolKeywordResultDto = sync.Pool{
	New: func() any {
		return new(KeywordResultDto)
	},
}

// GetKeywordResultDto() 从对象池中获取KeywordResultDto
func GetKeywordResultDto() *KeywordResultDto {
	return poolKeywordResultDto.Get().(*KeywordResultDto)
}

// ReleaseKeywordResultDto 释放KeywordResultDto
func ReleaseKeywordResultDto(v *KeywordResultDto) {
	v.TagList = v.TagList[:0]
	v.BasePrice = ""
	v.BuyCount = ""
	v.ClickCnt = ""
	v.ClickCostAvg = ""
	v.Cost = ""
	v.Ctr = ""
	v.ImpressionCnt = ""
	v.OnlineTime = ""
	v.Price = ""
	v.SearchCount = ""
	v.Status = ""
	v.Word = ""
	v.MatchCount = 0
	v.QsStar = 0
	v.Id = 0
	poolKeywordResultDto.Put(v)
}
