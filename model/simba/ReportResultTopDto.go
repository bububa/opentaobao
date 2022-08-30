package simba

// ReportResultTopDto 结构体
type ReportResultTopDto struct {
	// 时间
	Thedate string `json:"thedate,omitempty" xml:"thedate,omitempty"`
	// 点击率
	CtrStr string `json:"ctr_str,omitempty" xml:"ctr_str,omitempty"`
	// 平均点击花费
	CpcStr string `json:"cpc_str,omitempty" xml:"cpc_str,omitempty"`
	// 宝贝收藏率
	FavItemTotalCoverageStr string `json:"fav_item_total_coverage_str,omitempty" xml:"fav_item_total_coverage_str,omitempty"`
	// 加购率
	CartTotalCoverageStr string `json:"cart_total_coverage_str,omitempty" xml:"cart_total_coverage_str,omitempty"`
	// 投入产出比
	RoiStr string `json:"roi_str,omitempty" xml:"roi_str,omitempty"`
	// 点击转化率
	CoverageStr string `json:"coverage_str,omitempty" xml:"coverage_str,omitempty"`
	// 直接点击转化率
	DirectTransactionShippingCoverageStr string `json:"direct_transaction_shipping_coverage_str,omitempty" xml:"direct_transaction_shipping_coverage_str,omitempty"`
	// 总成交金额（单位元）
	TransactionTotalInYuanStr string `json:"transaction_total_in_yuan_str,omitempty" xml:"transaction_total_in_yuan_str,omitempty"`
	// 千次展现花费（单位元）
	CpmInYuanStr string `json:"cpm_in_yuan_str,omitempty" xml:"cpm_in_yuan_str,omitempty"`
	// 间接预售成交金额（单位元）
	IndirEprePayAmtInYuanStr string `json:"indir_epre_pay_amt_in_yuan_str,omitempty" xml:"indir_epre_pay_amt_in_yuan_str,omitempty"`
	// 平均点击花费（单位元）
	CpcInYuanStr string `json:"cpc_in_yuan_str,omitempty" xml:"cpc_in_yuan_str,omitempty"`
	// 直接预售成交金额（单位元）
	DirEprePayAmtInYuanStr string `json:"dir_epre_pay_amt_in_yuan_str,omitempty" xml:"dir_epre_pay_amt_in_yuan_str,omitempty"`
	// 购物金充值金额（单位元）
	ClickShoppingAmtInYuanStr string `json:"click_shopping_amt_in_yuan_str,omitempty" xml:"click_shopping_amt_in_yuan_str,omitempty"`
	// 尾款金额（单位元）
	HfhYsAmtInYuanStr string `json:"hfh_ys_amt_in_yuan_str,omitempty" xml:"hfh_ys_amt_in_yuan_str,omitempty"`
	// 加购成本（单位元）
	CartTotalCostInYuanStr string `json:"cart_total_cost_in_yuan_str,omitempty" xml:"cart_total_cost_in_yuan_str,omitempty"`
	// 直接成交金额（单位元）
	DirectTransactionInYuanStr string `json:"direct_transaction_in_yuan_str,omitempty" xml:"direct_transaction_in_yuan_str,omitempty"`
	// 间接成交金额（单位元）
	IndirectTransactionInYuanStr string `json:"indirect_transaction_in_yuan_str,omitempty" xml:"indirect_transaction_in_yuan_str,omitempty"`
	// 宝贝收藏成本（单位元）
	FavItemTotalCostInYuanStr string `json:"fav_item_total_cost_in_yuan_str,omitempty" xml:"fav_item_total_cost_in_yuan_str,omitempty"`
	// 总预售成交金额（单位元）
	EprePayAmtInYuanStr string `json:"epre_pay_amt_in_yuan_str,omitempty" xml:"epre_pay_amt_in_yuan_str,omitempty"`
	// 一口价金额（单位元）
	HfhYkjAmtInYuanStr string `json:"hfh_ykj_amt_in_yuan_str,omitempty" xml:"hfh_ykj_amt_in_yuan_str,omitempty"`
	// 花费（单位元）
	CostInYuanStr string `json:"cost_in_yuan_str,omitempty" xml:"cost_in_yuan_str,omitempty"`
	// 自然流量转化金额（单位元）
	SearchTransactionInYuanStr string `json:"search_transaction_in_yuan_str,omitempty" xml:"search_transaction_in_yuan_str,omitempty"`
	// 特权订金金额（单位元）
	HfhDjAmtInStr string `json:"hfh_dj_amt_in_str,omitempty" xml:"hfh_dj_amt_in_str,omitempty"`
	// 特权订金金额（单位元）
	HfhDjAmtInYuanStr string `json:"hfh_dj_amt_in_yuan_str,omitempty" xml:"hfh_dj_amt_in_yuan_str,omitempty"`
	// 平均展现排名
	AvgRankStr string `json:"avg_rank_str,omitempty" xml:"avg_rank_str,omitempty"`
	// 计划id
	CampaignId string `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 计划标题
	CampaignTitle string `json:"campaign_title,omitempty" xml:"campaign_title,omitempty"`
	// 计划类型名称
	CampaignTypeName string `json:"campaign_type_name,omitempty" xml:"campaign_type_name,omitempty"`
	// 单元标题
	AdgroupTitle string `json:"adgroup_title,omitempty" xml:"adgroup_title,omitempty"`
	// 跳转链接（商品详情页）
	Linkurl string `json:"linkurl,omitempty" xml:"linkurl,omitempty"`
	// 创意图片url
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 单元图片url
	AdgroupImgUrl string `json:"adgroup_img_url,omitempty" xml:"adgroup_img_url,omitempty"`
	// 创意标题
	CreativeTitle string `json:"creative_title,omitempty" xml:"creative_title,omitempty"`
	// 创意类型(1:静态创意，2：智能创意，3：PSD创意，4：动态文案创意，5：自动化创意，6：智能创意，7：综合静态创意，8：综合智能创意，9：自动直播创意，10：元素化创意，11：素材类创意，12：钻展多图创意)
	Creativetype string `json:"creativetype,omitempty" xml:"creativetype,omitempty"`
	// 关键词名称
	BidwordStr string `json:"bidword_str,omitempty" xml:"bidword_str,omitempty"`
	// 展现量
	Impression int64 `json:"impression,omitempty" xml:"impression,omitempty"`
	// 点击量
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 花费
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// 点击率
	Ctr int64 `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 平均点击花费
	Cpc int64 `json:"cpc,omitempty" xml:"cpc,omitempty"`
	// 千次展现花费
	Cpm int64 `json:"cpm,omitempty" xml:"cpm,omitempty"`
	// 总收藏数
	FavTotal int64 `json:"fav_total,omitempty" xml:"fav_total,omitempty"`
	// 收藏宝贝数
	FavItemTotal int64 `json:"fav_item_total,omitempty" xml:"fav_item_total,omitempty"`
	// 收藏店铺数
	FavShopTotal int64 `json:"fav_shop_total,omitempty" xml:"fav_shop_total,omitempty"`
	// 总购物车数
	CartTotal int64 `json:"cart_total,omitempty" xml:"cart_total,omitempty"`
	// 直接购物车数
	DirectCartTotal int64 `json:"direct_cart_total,omitempty" xml:"direct_cart_total,omitempty"`
	// 间接购物车数
	IndirectCartTotal int64 `json:"indirect_cart_total,omitempty" xml:"indirect_cart_total,omitempty"`
	// 加购成本
	CartTotalCost int64 `json:"cart_total_cost,omitempty" xml:"cart_total_cost,omitempty"`
	// 宝贝收藏成本
	FavItemTotalCost int64 `json:"fav_item_total_cost,omitempty" xml:"fav_item_total_cost,omitempty"`
	// 宝贝收藏率
	FavItemTotalCoverage int64 `json:"fav_item_total_coverage,omitempty" xml:"fav_item_total_coverage,omitempty"`
	// 加购率
	CartTotalCoverage int64 `json:"cart_total_coverage,omitempty" xml:"cart_total_coverage,omitempty"`
	// 总预售成交金额
	EprePayAmt int64 `json:"epre_pay_amt,omitempty" xml:"epre_pay_amt,omitempty"`
	// 总预售成交笔数
	EprePayCnt int64 `json:"epre_pay_cnt,omitempty" xml:"epre_pay_cnt,omitempty"`
	// 直接预售成交金额
	DirEprePayAmt int64 `json:"dir_epre_pay_amt,omitempty" xml:"dir_epre_pay_amt,omitempty"`
	// 直接预售成交笔数
	DirEprePayCnt int64 `json:"dir_epre_pay_cnt,omitempty" xml:"dir_epre_pay_cnt,omitempty"`
	// 间接预售成交金额
	IndirEprePayAmt int64 `json:"indir_epre_pay_amt,omitempty" xml:"indir_epre_pay_amt,omitempty"`
	// 间接预售成交笔数
	IndirEprePayCnt int64 `json:"indir_epre_pay_cnt,omitempty" xml:"indir_epre_pay_cnt,omitempty"`
	// 总成交金额
	TransactionTotal int64 `json:"transaction_total,omitempty" xml:"transaction_total,omitempty"`
	// 直接成交金额
	DirectTransaction int64 `json:"direct_transaction,omitempty" xml:"direct_transaction,omitempty"`
	// 间接成交金额
	IndirectTransaction int64 `json:"indirect_transaction,omitempty" xml:"indirect_transaction,omitempty"`
	// 总成交笔数
	TransactionShippingTotal int64 `json:"transaction_shipping_total,omitempty" xml:"transaction_shipping_total,omitempty"`
	// 直接成交笔数
	DirectTransactionShipping int64 `json:"direct_transaction_shipping,omitempty" xml:"direct_transaction_shipping,omitempty"`
	// 间接成交笔数
	IndirectTransactionShipping int64 `json:"indirect_transaction_shipping,omitempty" xml:"indirect_transaction_shipping,omitempty"`
	// 投入产出比
	Roi int64 `json:"roi,omitempty" xml:"roi,omitempty"`
	// 点击转化率
	Coverage int64 `json:"coverage,omitempty" xml:"coverage,omitempty"`
	// 直接点击转化率
	DirectTransactionShippingCoverage int64 `json:"direct_transaction_shipping_coverage,omitempty" xml:"direct_transaction_shipping_coverage,omitempty"`
	// 购物金充值笔数
	ClickShoppingNum int64 `json:"click_shopping_num,omitempty" xml:"click_shopping_num,omitempty"`
	// 购物金充值金额
	ClickShoppingAmt int64 `json:"click_shopping_amt,omitempty" xml:"click_shopping_amt,omitempty"`
	// 自然流量曝光
	SearchImpression int64 `json:"search_impression,omitempty" xml:"search_impression,omitempty"`
	// 自然流量转化金额
	SearchTransaction int64 `json:"search_transaction,omitempty" xml:"search_transaction,omitempty"`
	// 旺旺咨询量
	WwCnt int64 `json:"ww_cnt,omitempty" xml:"ww_cnt,omitempty"`
	// 特权订金订单支付笔数
	HfhDjCnt int64 `json:"hfh_dj_cnt,omitempty" xml:"hfh_dj_cnt,omitempty"`
	// 特权订金金额
	HfhDjAmt int64 `json:"hfh_dj_amt,omitempty" xml:"hfh_dj_amt,omitempty"`
	// 尾款支付笔数
	HfhYsCnt int64 `json:"hfh_ys_cnt,omitempty" xml:"hfh_ys_cnt,omitempty"`
	// 尾款金额
	HfhYsAmt int64 `json:"hfh_ys_amt,omitempty" xml:"hfh_ys_amt,omitempty"`
	// 一口价支付笔数
	HfhYkjCnt int64 `json:"hfh_ykj_cnt,omitempty" xml:"hfh_ykj_cnt,omitempty"`
	// 一口价金额
	HfhYkjAmt int64 `json:"hfh_ykj_amt,omitempty" xml:"hfh_ykj_amt,omitempty"`
	// 入会量
	RhCnt int64 `json:"rh_cnt,omitempty" xml:"rh_cnt,omitempty"`
	// 留资量
	LzCnt int64 `json:"lz_cnt,omitempty" xml:"lz_cnt,omitempty"`
	// 总成交金额（单位元）
	TransactionTotalInYuan int64 `json:"transaction_total_in_yuan,omitempty" xml:"transaction_total_in_yuan,omitempty"`
	// 千次展现花费（单位元）
	CpmInYuan int64 `json:"cpm_in_yuan,omitempty" xml:"cpm_in_yuan,omitempty"`
	// 间接预售成交金额（单位元）
	IndirEprePayAmtInYuan int64 `json:"indir_epre_pay_amt_in_yuan,omitempty" xml:"indir_epre_pay_amt_in_yuan,omitempty"`
	// 平均点击花费（单位元）
	CpcInYuan int64 `json:"cpc_in_yuan,omitempty" xml:"cpc_in_yuan,omitempty"`
	// 直接预售成交金额（单位元）
	DirEprePayAmtInYuan int64 `json:"dir_epre_pay_amt_in_yuan,omitempty" xml:"dir_epre_pay_amt_in_yuan,omitempty"`
	// 购物金充值金额（单位元）
	ClickShoppingAmtInYuan int64 `json:"click_shopping_amt_in_yuan,omitempty" xml:"click_shopping_amt_in_yuan,omitempty"`
	// 尾款金额（单位元）
	HfhYsAmtInYuan int64 `json:"hfh_ys_amt_in_yuan,omitempty" xml:"hfh_ys_amt_in_yuan,omitempty"`
	// 加购成本（单位元）
	CartTotalCostInYuan int64 `json:"cart_total_cost_in_yuan,omitempty" xml:"cart_total_cost_in_yuan,omitempty"`
	// 直接成交金额（单位元）
	DirectTransactionInYuan int64 `json:"direct_transaction_in_yuan,omitempty" xml:"direct_transaction_in_yuan,omitempty"`
	// 间接成交金额（单位元）
	IndirectTransactionInYuan int64 `json:"indirect_transaction_in_yuan,omitempty" xml:"indirect_transaction_in_yuan,omitempty"`
	// 宝贝收藏成本（单位元）
	FavItemTotalCostInYuan int64 `json:"fav_item_total_cost_in_yuan,omitempty" xml:"fav_item_total_cost_in_yuan,omitempty"`
	// 总预售成交金额（单位元）
	EprePayAmtInYuan int64 `json:"epre_pay_amt_in_yuan,omitempty" xml:"epre_pay_amt_in_yuan,omitempty"`
	// 一口价金额（单位元）
	HfhYkjAmtInYuan int64 `json:"hfh_ykj_amt_in_yuan,omitempty" xml:"hfh_ykj_amt_in_yuan,omitempty"`
	// 花费（单位元）
	CostInYuan int64 `json:"cost_in_yuan,omitempty" xml:"cost_in_yuan,omitempty"`
	// 自然流量转化金额（单位元）
	SearchTransactionInYuan int64 `json:"search_transaction_in_yuan,omitempty" xml:"search_transaction_in_yuan,omitempty"`
	// 特权订金金额（单位元）
	HfhDjAmtInYuan int64 `json:"hfh_dj_amt_in_yuan,omitempty" xml:"hfh_dj_amt_in_yuan,omitempty"`
	// 计划类型（直通车搜索-无线/pc：0；智能推广计划：8；销量明星计划：16；口碑L店计划：17；新享一键推广计划-独立结算账户(策略中心)：21；超级直播-一键推广计划（策略中心：订单模式、计划不复用：22；大快消一键推广计划（策略中心）：23；超级直播-持续推广计划（策略中心:计划模式、可复用、类似单品）：24；合约广告、流量卡计划：31；极速推计划：37；AI智投：38；）
	CampaignType int64 `json:"campaign_type,omitempty" xml:"campaign_type,omitempty"`
	// 单元id
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 特权订金金额（单位元）
	HfhDjAmtIn int64 `json:"hfh_dj_amt_in,omitempty" xml:"hfh_dj_amt_in,omitempty"`
	// 人群id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 创意id
	Creativeid int64 `json:"creativeid,omitempty" xml:"creativeid,omitempty"`
	// 无线出价
	WirelessPrice int64 `json:"wireless_price,omitempty" xml:"wireless_price,omitempty"`
	// 平均展现排名
	AvgRank int64 `json:"avg_rank,omitempty" xml:"avg_rank,omitempty"`
	// pc出价
	PcPrice int64 `json:"pc_price,omitempty" xml:"pc_price,omitempty"`
	// 日限额
	CampaignBudget int64 `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`
	// 关键词ID
	BidwordId int64 `json:"bidword_id,omitempty" xml:"bidword_id,omitempty"`
}
