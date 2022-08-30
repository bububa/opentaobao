package tbk

// TopNInfoDto 结构体
type TopNInfoDto struct {
	// 百亿补贴商品特征标签，eg.今日发货、晚发补偿、限购一件等
	BybtItemTags []string `json:"bybt_item_tags,omitempty" xml:"bybt_item_tags>string,omitempty"`
	// 前N件佣金结束时间
	TopnEndTime string `json:"topn_end_time,omitempty" xml:"topn_end_time,omitempty"`
	// 前N件佣金开始时间
	TopnStartTime string `json:"topn_start_time,omitempty" xml:"topn_start_time,omitempty"`
	// 前N件佣金率
	TopnRate string `json:"topn_rate,omitempty" xml:"topn_rate,omitempty"`
	// 百亿补贴品牌logo
	BybtBrandLogo string `json:"bybt_brand_logo,omitempty" xml:"bybt_brand_logo,omitempty"`
	// 百亿补贴白底图
	BybtPicUrl string `json:"bybt_pic_url,omitempty" xml:"bybt_pic_url,omitempty"`
	// 百亿补贴专属券面额，仅限百亿补贴场景透出
	BybtCouponAmount string `json:"bybt_coupon_amount,omitempty" xml:"bybt_coupon_amount,omitempty"`
	// 百亿补贴页面实时价
	BybtShowPrice string `json:"bybt_show_price,omitempty" xml:"bybt_show_price,omitempty"`
	// 全网对比参考价格
	BybtLowestPrice string `json:"bybt_lowest_price,omitempty" xml:"bybt_lowest_price,omitempty"`
	// 商品的百亿补贴开始时间
	BybtEndTime string `json:"bybt_end_time,omitempty" xml:"bybt_end_time,omitempty"`
	// 商品的百亿补贴结束时间
	BybtStartTime string `json:"bybt_start_time,omitempty" xml:"bybt_start_time,omitempty"`
	// 前N件剩余库存
	TopnQuantity int64 `json:"topn_quantity,omitempty" xml:"topn_quantity,omitempty"`
	// 前N件初始总库存
	TopnTotalCount int64 `json:"topn_total_count,omitempty" xml:"topn_total_count,omitempty"`
}
