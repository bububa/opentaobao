package icbuseller

// VendorMerchantRecordBaseDto 结构体
type VendorMerchantRecordBaseDto struct {
	// 定制力是否双60
	Nrts6060 string `json:"nrts6060,omitempty" xml:"nrts6060,omitempty"`
	// 赛道，RTS/询盘
	Road string `json:"road,omitempty" xml:"road,omitempty"`
	// 是否至少发布了1个良好的且非aliwood且非段灯片的视频
	IsVideo string `json:"is_video,omitempty" xml:"is_video,omitempty"`
	// 是否至少发布了1个有效报价RFQ
	IsRfq string `json:"is_rfq,omitempty" xml:"is_rfq,omitempty"`
	// 历史截至当日橱窗商品使用率
	WinProdRatioStd001 string `json:"win_prod_ratio_std001,omitempty" xml:"win_prod_ratio_std001,omitempty"`
	// 是否信保亮灯Y/N
	IsCrdScrty string `json:"is_crd_scrty,omitempty" xml:"is_crd_scrty,omitempty"`
	// 月PC在线时长达标率
	PcOnlineHourRate string `json:"pc_online_hour_rate,omitempty" xml:"pc_online_hour_rate,omitempty"`
	// 近30天PC/MA在线时长(小时)
	StayHourPcM string `json:"stay_hour_pc_m,omitempty" xml:"stay_hour_pc_m,omitempty"`
	// 月pc在线天数达标率
	PcOnlineDayRate string `json:"pc_online_day_rate,omitempty" xml:"pc_online_day_rate,omitempty"`
	// 月活跃子账号达标率
	ActiveMbrRate string `json:"active_mbr_rate,omitempty" xml:"active_mbr_rate,omitempty"`
	// 月mc达标率
	McRate string `json:"mc_rate,omitempty" xml:"mc_rate,omitempty"`
	// 关键词达标率
	KwRate string `json:"kw_rate,omitempty" xml:"kw_rate,omitempty"`
	// k200比率，关键词除以200的比率
	Kw200Rate string `json:"kw200_rate,omitempty" xml:"kw200_rate,omitempty"`
	// 是否有侵权品牌Y/N
	IsTortBrand string `json:"is_tort_brand,omitempty" xml:"is_tort_brand,omitempty"`
	// 是否有价格不合理产品Y/N
	HavePriceProd string `json:"have_price_prod,omitempty" xml:"have_price_prod,omitempty"`
	// 是否有重复铺货产品Y/N
	HaveRepeatProd string `json:"have_repeat_prod,omitempty" xml:"have_repeat_prod,omitempty"`
	// 实力优品达标率
	GoodProdRate string `json:"good_prod_rate,omitempty" xml:"good_prod_rate,omitempty"`
	// 潜力品达标率
	PotentialProdRate string `json:"potential_prod_rate,omitempty" xml:"potential_prod_rate,omitempty"`
	// p150且4分达标率
	P1504Rate string `json:"p1504_rate,omitempty" xml:"p1504_rate,omitempty"`
	// p30且4分达标率
	P304Rate string `json:"p304_rate,omitempty" xml:"p304_rate,omitempty"`
	// 店铺是否装修，页面最晚修改时间和旺铺创建时间差
	IsCompDec string `json:"is_comp_dec,omitempty" xml:"is_comp_dec,omitempty"`
	// 订单完成时间
	TradesuccessTime string `json:"tradesuccess_time,omitempty" xml:"tradesuccess_time,omitempty"`
	// 进入交付时间
	DeliveredTime string `json:"delivered_time,omitempty" xml:"delivered_time,omitempty"`
	// 下单时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 交易单当前状态
	CurrentStatus string `json:"current_status,omitempty" xml:"current_status,omitempty"`
	// 订单号
	OrdNum string `json:"ord_num,omitempty" xml:"ord_num,omitempty"`
	// 服务名称
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// 服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务商编码
	VendorCode string `json:"vendor_code,omitempty" xml:"vendor_code,omitempty"`
	// 服务商公司名
	VendorCompName string `json:"vendor_comp_name,omitempty" xml:"vendor_comp_name,omitempty"`
	// 主营二级类目描述
	MainCateLv2Desc string `json:"main_cate_lv2_desc,omitempty" xml:"main_cate_lv2_desc,omitempty"`
	// 主营一级类目描述
	MainCateLv1Desc string `json:"main_cate_lv1_desc,omitempty" xml:"main_cate_lv1_desc,omitempty"`
	// 会员主账号id
	AdminMbrId string `json:"admin_mbr_id,omitempty" xml:"admin_mbr_id,omitempty"`
	// 系统入库时间
	DwInsTime string `json:"dw_ins_time,omitempty" xml:"dw_ins_time,omitempty"`
	// 统计日期
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// RTS力是否双60
	Rts6060 string `json:"rts6060,omitempty" xml:"rts6060,omitempty"`
	// 最近7天店铺日均搜索曝光是否正常Y/N
	IsCompExpOk string `json:"is_comp_exp_ok,omitempty" xml:"is_comp_exp_ok,omitempty"`
	// 关键词曝光是否正常Y/N
	IsKwExpOk string `json:"is_kw_exp_ok,omitempty" xml:"is_kw_exp_ok,omitempty"`
	// 是否有降星风险Y/N
	IsStarLowerRisk string `json:"is_star_lower_risk,omitempty" xml:"is_star_lower_risk,omitempty"`
	// 是否p4p月消耗达标【Y/N】
	IsP4pCostQualified1 string `json:"is_p4p_cost_qualified1,omitempty" xml:"is_p4p_cost_qualified1,omitempty"`
	// 是否p4p绑定推广商品数大于等于80【Y/N】
	IsP4pProdGeq80Cnt string `json:"is_p4p_prod_geq80_cnt,omitempty" xml:"is_p4p_prod_geq80_cnt,omitempty"`
	// 是否p4p绑定推广商品数大于等于30【Y/N】
	IsP4pProdGeq30Cnt string `json:"is_p4p_prod_geq30_cnt,omitempty" xml:"is_p4p_prod_geq30_cnt,omitempty"`
	// 是否商机破蛋【Y/N】
	IsHasAb string `json:"is_has_ab,omitempty" xml:"is_has_ab,omitempty"`
	// 是否首月询盘破蛋【Y/N】
	IsHasMcFstMon string `json:"is_has_mc_fst_mon,omitempty" xml:"is_has_mc_fst_mon,omitempty"`
	// 是否MC15达标（询盘达到15个）
	IsMc15 string `json:"is_mc15,omitempty" xml:"is_mc15,omitempty"`
	// 服务商开始接洽时间
	ContactTime string `json:"contact_time,omitempty" xml:"contact_time,omitempty"`
	// 是否蓝海定招
	IsBlueNewSign string `json:"is_blue_new_sign,omitempty" xml:"is_blue_new_sign,omitempty"`
	// 商机是否达到首年客户二级行业均值【Y/N】（近30天）
	IsReachNewCustAbCatelv2 string `json:"is_reach_new_cust_ab_catelv2,omitempty" xml:"is_reach_new_cust_ab_catelv2,omitempty"`
	// 商家平台开通首半年P4P月活（月消耗200元现金以上）达标次数
	P4pActivityMonths string `json:"p4p_activity_months,omitempty" xml:"p4p_activity_months,omitempty"`
	// 近15天内P30潜力品是否曾达标【Y/N】
	IsP30PotentialProdRate15d string `json:"is_p30_potential_prod_rate15d,omitempty" xml:"is_p30_potential_prod_rate15d,omitempty"`
	// P30潜力品达标率(潜力品数量/30)
	P30PotentialProdRate string `json:"p30_potential_prod_rate,omitempty" xml:"p30_potential_prod_rate,omitempty"`
	// 近30天内P150潜力品是否曾达标【Y/N】
	IsP150PotentialProdRate30d string `json:"is_p150_potential_prod_rate30d,omitempty" xml:"is_p150_potential_prod_rate30d,omitempty"`
	// 近60天内P150潜力品是否曾达标【Y/N】
	IsP150PotentialProdRate60d string `json:"is_p150_potential_prod_rate60d,omitempty" xml:"is_p150_potential_prod_rate60d,omitempty"`
	// P150潜力品达标率(潜力品数量/150)
	P150PotentialProdRate string `json:"p150_potential_prod_rate,omitempty" xml:"p150_potential_prod_rate,omitempty"`
	// 近60天内P200潜力品是否曾达标【Y/N】
	IsP200PotentialProdRate60d string `json:"is_p200_potential_prod_rate60d,omitempty" xml:"is_p200_potential_prod_rate60d,omitempty"`
	// P200潜力品达标率(潜力品数量/200)
	P200PotentialProdRate string `json:"p200_potential_prod_rate,omitempty" xml:"p200_potential_prod_rate,omitempty"`
	// 近90天内P300潜力品是否曾达标【Y/N】
	IsP300PotentialProdRate90d string `json:"is_p300_potential_prod_rate90d,omitempty" xml:"is_p300_potential_prod_rate90d,omitempty"`
	// P300潜力品达标率(潜力品数量/300)
	P300PotentialProdRate string `json:"p300_potential_prod_rate,omitempty" xml:"p300_potential_prod_rate,omitempty"`
	// 商机是否达二级行业均值【Y/N】
	IsReachAbCatelv2 string `json:"is_reach_ab_catelv2,omitempty" xml:"is_reach_ab_catelv2,omitempty"`
	// 当前cgs合同的服务开始时间
	CrntGsActualBeginTime string `json:"crnt_gs_actual_begin_time,omitempty" xml:"crnt_gs_actual_begin_time,omitempty"`
	// 当前月份服务期间总财务消耗
	P4pCostAmtCurrentmon string `json:"p4p_cost_amt_currentmon,omitempty" xml:"p4p_cost_amt_currentmon,omitempty"`
	// 授权前30天二级行业CTR是行业均值的X
	P4pCatelv2Ctr1Befer30Percent string `json:"p4p_catelv2_ctr1_befer30_percent,omitempty" xml:"p4p_catelv2_ctr1_befer30_percent,omitempty"`
	// 授权服务当月客户二级行业CTR是行业均值的X
	P4pCatelv2Ctr1CurrentmonPercent string `json:"p4p_catelv2_ctr1_currentmon_percent,omitempty" xml:"p4p_catelv2_ctr1_currentmon_percent,omitempty"`
	// 当月托管服务是否二级行业CTR达标
	P4pIspassCurrentmon string `json:"p4p_ispass_currentmon,omitempty" xml:"p4p_ispass_currentmon,omitempty"`
	// 当月代操服务佣金预估
	P4pExpectFeeCurrentmon string `json:"p4p_expect_fee_currentmon,omitempty" xml:"p4p_expect_fee_currentmon,omitempty"`
	// 订单接洽时间
	ContactDate string `json:"contact_date,omitempty" xml:"contact_date,omitempty"`
	// 托管开始日期
	P4pHostingServiceBeginTime string `json:"p4p_hosting_service_begin_time,omitempty" xml:"p4p_hosting_service_begin_time,omitempty"`
	// 托管结束日期
	P4pHostingServiceEndTime string `json:"p4p_hosting_service_end_time,omitempty" xml:"p4p_hosting_service_end_time,omitempty"`
	// 服务天数
	P4pFuwuDays string `json:"p4p_fuwu_days,omitempty" xml:"p4p_fuwu_days,omitempty"`
	// 服务类型
	P4pServiceType string `json:"p4p_service_type,omitempty" xml:"p4p_service_type,omitempty"`
	// 是否服务中
	P4pIsSrvc string `json:"p4p_is_srvc,omitempty" xml:"p4p_is_srvc,omitempty"`
	// 当前账户余额
	P4pTotalBalance string `json:"p4p_total_balance,omitempty" xml:"p4p_total_balance,omitempty"`
	// 服务前L层级
	P4pLevelBefore string `json:"p4p_level_before,omitempty" xml:"p4p_level_before,omitempty"`
	// 当月L层级
	P4pLevelCurrent string `json:"p4p_level_current,omitempty" xml:"p4p_level_current,omitempty"`
	// 客户服务前30操作天数
	P4pOptDaysBefer30 string `json:"p4p_opt_days_befer30,omitempty" xml:"p4p_opt_days_befer30,omitempty"`
	// 当前月份服务期间代理商操作天数
	P4pAgencyOptDaysCurrentmon string `json:"p4p_agency_opt_days_currentmon,omitempty" xml:"p4p_agency_opt_days_currentmon,omitempty"`
	// 服务前30天推广天数
	P4pOnDaysBefer30 string `json:"p4p_on_days_befer30,omitempty" xml:"p4p_on_days_befer30,omitempty"`
	// 服务前30天消耗天数
	P4pCostDaysBefer30 string `json:"p4p_cost_days_befer30,omitempty" xml:"p4p_cost_days_befer30,omitempty"`
	// 服务前30天有推广的日均曝光是二级行业均值的X
	P4pCatelv2AvgdailyImprCntBefer30Percent string `json:"p4p_catelv2_avgdaily_impr_cnt_befer30_percent,omitempty" xml:"p4p_catelv2_avgdaily_impr_cnt_befer30_percent,omitempty"`
	// 服务前30天有推广的日均点击是行业均值的X
	P4pCatelv2AvgdailyClickCntBefer30Percent string `json:"p4p_catelv2_avgdaily_click_cnt_befer30_percent,omitempty" xml:"p4p_catelv2_avgdaily_click_cnt_befer30_percent,omitempty"`
	// 服务前30天日均推广关键词数
	P4pAvgKwordBefer30 string `json:"p4p_avg_kword_befer30,omitempty" xml:"p4p_avg_kword_befer30,omitempty"`
	// 服务前30天日均推广品数
	P4pAvgdailyOnprodsBefer30 string `json:"p4p_avgdaily_onprods_befer30,omitempty" xml:"p4p_avgdaily_onprods_befer30,omitempty"`
	// 服务至今推广天数
	P4pOnDaysService2now string `json:"p4p_on_days_service2now,omitempty" xml:"p4p_on_days_service2now,omitempty"`
	// 当前月份服务期间推广天数
	P4pOnDaysCurrentmon string `json:"p4p_on_days_currentmon,omitempty" xml:"p4p_on_days_currentmon,omitempty"`
	// 服务至今消耗天数
	P4pCostDaysService2now string `json:"p4p_cost_days_service2now,omitempty" xml:"p4p_cost_days_service2now,omitempty"`
	// 当前月份服务期间消耗天数
	P4pCostDaysCurrentmon string `json:"p4p_cost_days_currentmon,omitempty" xml:"p4p_cost_days_currentmon,omitempty"`
	// 服务至今总财务消耗
	P4pCostAmtService2now string `json:"p4p_cost_amt_service2now,omitempty" xml:"p4p_cost_amt_service2now,omitempty"`
	// 当前月份服务期间日均推广关键词数
	P4pAvgKwordCurrentmon string `json:"p4p_avg_kword_currentmon,omitempty" xml:"p4p_avg_kword_currentmon,omitempty"`
	// 服务期间日均推广品数
	P4pAvgdailyOnprodsCurrentmon string `json:"p4p_avgdaily_onprods_currentmon,omitempty" xml:"p4p_avgdaily_onprods_currentmon,omitempty"`
	// 服务前30天CPC是二级行业均值X%
	P4pCatelv2Cpc1Befer30Percent string `json:"p4p_catelv2_cpc1_befer30_percent,omitempty" xml:"p4p_catelv2_cpc1_befer30_percent,omitempty"`
	// 服务前30天日均P4P预算
	P4pAvgdailyBudgetBefer30 string `json:"p4p_avgdaily_budget_befer30,omitempty" xml:"p4p_avgdaily_budget_befer30,omitempty"`
	// 服务前30天总财务消耗
	P4pAvgdailyCostAmtBefer30 string `json:"p4p_avgdaily_cost_amt_befer30,omitempty" xml:"p4p_avgdaily_cost_amt_befer30,omitempty"`
	// 当前月份服务期间日均P4P预算
	P4pAvgdailyBudgetCurrentmon string `json:"p4p_avgdaily_budget_currentmon,omitempty" xml:"p4p_avgdaily_budget_currentmon,omitempty"`
	// 当前月份服务期间日均P4P财务消耗
	P4pAvgdailyCostAmtCurrentmon string `json:"p4p_avgdaily_cost_amt_currentmon,omitempty" xml:"p4p_avgdaily_cost_amt_currentmon,omitempty"`
	// 服务期间日均曝光是二级行业均值的X（当前月份服务期间日均P4P曝光/(当月服务状态近30天大盘二级客均曝光/30)）
	P4pCatelv2AvgdailyImprCntCurrentmonPercent string `json:"p4p_catelv2_avgdaily_impr_cnt_currentmon_percent,omitempty" xml:"p4p_catelv2_avgdaily_impr_cnt_currentmon_percent,omitempty"`
	// 服务期间日均点击是二级行业均值X
	P4pCatelv2AvgdailyClickCntCurrentmon string `json:"p4p_catelv2_avgdaily_click_cnt_currentmon,omitempty" xml:"p4p_catelv2_avgdaily_click_cnt_currentmon,omitempty"`
	// AB30
	AvgAb30 string `json:"avg_ab30,omitempty" xml:"avg_ab30,omitempty"`
	// 平台开通15天内的新手任务是否完成
	NewtastStat string `json:"newtast_stat,omitempty" xml:"newtast_stat,omitempty"`
	// 平台开通30天内的P4P是否开通
	IsP4pOn30d string `json:"is_p4p_on_30d,omitempty" xml:"is_p4p_on_30d,omitempty"`
	// 实力优品中参与P4P推广的比例
	OnExcellentProdRate string `json:"on_excellent_prod_rate,omitempty" xml:"on_excellent_prod_rate,omitempty"`
	// 下单时时P4PL层级
	P4pStarOrd string `json:"p4p_star_ord,omitempty" xml:"p4p_star_ord,omitempty"`
	// P4P当日是否推广中Y/N
	P4pIsOn string `json:"p4p_is_on,omitempty" xml:"p4p_is_on,omitempty"`
	// 是否曾经是优化师库内商家Y/N
	P4pHadOptSrvc string `json:"p4p_had_opt_srvc,omitempty" xml:"p4p_had_opt_srvc,omitempty"`
	// 服务期间客户的CTR是服务前30天客户CTR的X
	P4pCtrService2nowTimes string `json:"p4p_ctr_service2now_times,omitempty" xml:"p4p_ctr_service2now_times,omitempty"`
	// 服务期间当月日均消耗是否超过优化师服务期间日均消耗Y/N
	P4pIsExceedAvgdailyCostCurrentmon string `json:"p4p_is_exceed_avgdaily_cost_currentmon,omitempty" xml:"p4p_is_exceed_avgdaily_cost_currentmon,omitempty"`
	// 服务期间CPF(mc+atm+ord)是服务前30天CPF的X
	P4pCpf2Service2nowTimes string `json:"p4p_cpf2_service2now_times,omitempty" xml:"p4p_cpf2_service2now_times,omitempty"`
	// 近30天PC/MA活跃天数
	ActiveDayPcM int64 `json:"active_day_pc_m,omitempty" xml:"active_day_pc_m,omitempty"`
	// 近30天活跃子账号数
	ActiveMbrM int64 `json:"active_mbr_m,omitempty" xml:"active_mbr_m,omitempty"`
	// 视频数量
	ProdVideoCnt int64 `json:"prod_video_cnt,omitempty" xml:"prod_video_cnt,omitempty"`
	// RTS商品数
	RtsProdCnt int64 `json:"rts_prod_cnt,omitempty" xml:"rts_prod_cnt,omitempty"`
	// 实力优品数
	GoodProdCnt int64 `json:"good_prod_cnt,omitempty" xml:"good_prod_cnt,omitempty"`
	// 潜力品数
	PotentialProdCnt int64 `json:"potential_prod_cnt,omitempty" xml:"potential_prod_cnt,omitempty"`
	// 蓝海品数
	ProdBlue int64 `json:"prod_blue,omitempty" xml:"prod_blue,omitempty"`
	// 4.5分及以上商品数
	Prod45 int64 `json:"prod45,omitempty" xml:"prod45,omitempty"`
	// 4.0分及以上商品数
	Prod40 int64 `json:"prod40,omitempty" xml:"prod40,omitempty"`
	// 有效商品数
	ProdCnt int64 `json:"prod_cnt,omitempty" xml:"prod_cnt,omitempty"`
	// 服务商发品数
	AgencyProdCnt int64 `json:"agency_prod_cnt,omitempty" xml:"agency_prod_cnt,omitempty"`
	// 主营二级类目ID
	MainCateLv2Std001 int64 `json:"main_cate_lv2_std001,omitempty" xml:"main_cate_lv2_std001,omitempty"`
	// 主营一级类目ID
	MainCateLv1Std001 int64 `json:"main_cate_lv1_std001,omitempty" xml:"main_cate_lv1_std001,omitempty"`
	// 当月星等级
	LevelStarPage int64 `json:"level_star_page,omitempty" xml:"level_star_page,omitempty"`
	// 预测下月星等级
	LevelStar int64 `json:"level_star,omitempty" xml:"level_star,omitempty"`
	// 开通周期（付款到开通）
	OpenCycle int64 `json:"open_cycle,omitempty" xml:"open_cycle,omitempty"`
	// 首年0-6个月信保渗透月数
	FiscalSvrc6mValidOrdCnt int64 `json:"fiscal_svrc6m_valid_ord_cnt,omitempty" xml:"fiscal_svrc6m_valid_ord_cnt,omitempty"`
	// 最近一次开通服务天数
	LatestActualSrvcDays int64 `json:"latest_actual_srvc_days,omitempty" xml:"latest_actual_srvc_days,omitempty"`
	// 商家主动评价分数
	ProdScore int64 `json:"prod_score,omitempty" xml:"prod_score,omitempty"`
	// 近30天mc完成率
	McComplianceRate30d int64 `json:"mc_compliance_rate30d,omitempty" xml:"mc_compliance_rate30d,omitempty"`
	// 近30天AB完成率
	AbComplianceRate30d int64 `json:"ab_compliance_rate30d,omitempty" xml:"ab_compliance_rate30d,omitempty"`
	// 商机数与二级行业均值的差距（即：未达标时距离目标值差值，达标则不显示值
	Catelv2AvgAbDiff int64 `json:"catelv2_avg_ab_diff,omitempty" xml:"catelv2_avg_ab_diff,omitempty"`
	// p4p层级1及以上连续月数
	P4pStarMonths6m int64 `json:"p4p_star_months_6m,omitempty" xml:"p4p_star_months_6m,omitempty"`
	// P4P连续活跃月份数
	P4pCostMonths6m int64 `json:"p4p_cost_months_6m,omitempty" xml:"p4p_cost_months_6m,omitempty"`
}
