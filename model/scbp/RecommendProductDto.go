package scbp

// RecommendProductDto 结构体
type RecommendProductDto struct {
	// 品标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 消耗
	AdsCostAmt string `json:"ads_cost_amt,omitempty" xml:"ads_cost_amt,omitempty"`
	// 点击率
	AdsCtr string `json:"ads_ctr,omitempty" xml:"ads_ctr,omitempty"`
	// 平均点击价格
	AdsCpc string `json:"ads_cpc,omitempty" xml:"ads_cpc,omitempty"`
	// 网页点击率
	WebCtr string `json:"web_ctr,omitempty" xml:"web_ctr,omitempty"`
	// 产品分数
	ProdScore string `json:"prod_score,omitempty" xml:"prod_score,omitempty"`
	// 权力分数
	PowerScore string `json:"power_score,omitempty" xml:"power_score,omitempty"`
	// 蓝海指数
	BlueSeaScore string `json:"blue_sea_score,omitempty" xml:"blue_sea_score,omitempty"`
	// 品创建时间
	GmtPostingCreate string `json:"gmt_posting_create,omitempty" xml:"gmt_posting_create,omitempty"`
	// 品id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 曝光数
	AdsImprCnt int64 `json:"ads_impr_cnt,omitempty" xml:"ads_impr_cnt,omitempty"`
	// 点击数
	AdsClickCnt int64 `json:"ads_click_cnt,omitempty" xml:"ads_click_cnt,omitempty"`
	// 网页曝光数
	WebImprCnt int64 `json:"web_impr_cnt,omitempty" xml:"web_impr_cnt,omitempty"`
	// 网页点击数
	WebClickCnt int64 `json:"web_click_cnt,omitempty" xml:"web_click_cnt,omitempty"`
	// 网页排名
	WebRank int64 `json:"web_rank,omitempty" xml:"web_rank,omitempty"`
	// 排名
	AdsRank int64 `json:"ads_rank,omitempty" xml:"ads_rank,omitempty"`
	// 商品推广状态
	PromotionStatus int64 `json:"promotion_status,omitempty" xml:"promotion_status,omitempty"`
	// 优越品
	SuperiorProd int64 `json:"superior_prod,omitempty" xml:"superior_prod,omitempty"`
}
