package scbp

import (
	"sync"
)

// KeywordRankPriceDto 结构体
type KeywordRankPriceDto struct {
	// 关键词前五名排价
	PriceArray []string `json:"price_array,omitempty" xml:"price_array>string,omitempty"`
	// 价格列表(计划)(不建议直接使用)
	PriceList []int64 `json:"price_list,omitempty" xml:"price_list>int64,omitempty"`
	// 价格列表(客户)(不建议直接使用)
	CustPriceList []int64 `json:"cust_price_list,omitempty" xml:"cust_price_list>int64,omitempty"`
	// 价格列表(客户)(元)(低价处理后结果)
	CustPriceArray []string `json:"cust_price_array,omitempty" xml:"cust_price_array>string,omitempty"`
	// 排价结果
	RankPriceList []string `json:"rank_price_list,omitempty" xml:"rank_price_list>string,omitempty"`
	// 关键词
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 公司ID
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
}

var poolKeywordRankPriceDto = sync.Pool{
	New: func() any {
		return new(KeywordRankPriceDto)
	},
}

// GetKeywordRankPriceDto() 从对象池中获取KeywordRankPriceDto
func GetKeywordRankPriceDto() *KeywordRankPriceDto {
	return poolKeywordRankPriceDto.Get().(*KeywordRankPriceDto)
}

// ReleaseKeywordRankPriceDto 释放KeywordRankPriceDto
func ReleaseKeywordRankPriceDto(v *KeywordRankPriceDto) {
	v.PriceArray = v.PriceArray[:0]
	v.PriceList = v.PriceList[:0]
	v.CustPriceList = v.CustPriceList[:0]
	v.CustPriceArray = v.CustPriceArray[:0]
	v.RankPriceList = v.RankPriceList[:0]
	v.Keyword = ""
	v.CompanyId = 0
	v.CampaignId = 0
	poolKeywordRankPriceDto.Put(v)
}
