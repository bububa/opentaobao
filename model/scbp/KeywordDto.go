package scbp

import (
	"sync"
)

// KeywordDto 结构体
type KeywordDto struct {
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 更新时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 词
	Word string `json:"word,omitempty" xml:"word,omitempty"`
	// 普通词
	NormWord string `json:"norm_word,omitempty" xml:"norm_word,omitempty"`
	// 出价分
	BidPrice string `json:"bid_price,omitempty" xml:"bid_price,omitempty"`
	// 平均出价
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 建议出价
	SugPrice string `json:"sug_price,omitempty" xml:"sug_price,omitempty"`
	// 低价
	BasePrice string `json:"base_price,omitempty" xml:"base_price,omitempty"`
	// 配置信息
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 效果数据
	Effect *KeywordEffectDto `json:"effect,omitempty" xml:"effect,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 产品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 计划id
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 状态
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// 星级
	QsStar int64 `json:"qs_star,omitempty" xml:"qs_star,omitempty"`
	// 最优商品id
	BestMatchProduct int64 `json:"best_match_product,omitempty" xml:"best_match_product,omitempty"`
	// 相关性产品数量
	RelativeProductsCount int64 `json:"relative_products_count,omitempty" xml:"relative_products_count,omitempty"`
	// 搜索数量
	SearchCount int64 `json:"search_count,omitempty" xml:"search_count,omitempty"`
	// 购买数量
	BuyCount int64 `json:"buy_count,omitempty" xml:"buy_count,omitempty"`
	// 战略数据
	BidStrategy *BidStrategyDto `json:"bid_strategy,omitempty" xml:"bid_strategy,omitempty"`
}

var poolKeywordDto = sync.Pool{
	New: func() any {
		return new(KeywordDto)
	},
}

// GetKeywordDto() 从对象池中获取KeywordDto
func GetKeywordDto() *KeywordDto {
	return poolKeywordDto.Get().(*KeywordDto)
}

// ReleaseKeywordDto 释放KeywordDto
func ReleaseKeywordDto(v *KeywordDto) {
	v.GmtCreate = ""
	v.GmtModify = ""
	v.Word = ""
	v.NormWord = ""
	v.BidPrice = ""
	v.AvgPrice = ""
	v.SugPrice = ""
	v.BasePrice = ""
	v.Properties = ""
	v.Effect = nil
	v.Id = 0
	v.ProductId = 0
	v.CampaignId = 0
	v.OnlineStatus = 0
	v.QsStar = 0
	v.BestMatchProduct = 0
	v.RelativeProductsCount = 0
	v.SearchCount = 0
	v.BuyCount = 0
	v.BidStrategy = nil
	poolKeywordDto.Put(v)
}
