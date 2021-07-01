package icbuseller

// VendorMerchantRecordBaseDto 
type VendorMerchantRecordBaseDto struct {
    // 定制力是否双60
    Nrts6060   string `json:"nrts6060,omitempty" xml:"nrts6060,omitempty"`
    // 赛道，RTS/询盘
    Road   string `json:"road,omitempty" xml:"road,omitempty"`
    // 是否至少发布了1个良好的且非aliwood且非段灯片的视频
    IsVideo   string `json:"is_video,omitempty" xml:"is_video,omitempty"`
    // 是否至少发布了1个有效报价RFQ
    IsRfq   string `json:"is_rfq,omitempty" xml:"is_rfq,omitempty"`
    // 历史截至当日橱窗商品使用率
    WinProdRatioStd001   string `json:"win_prod_ratio_std001,omitempty" xml:"win_prod_ratio_std001,omitempty"`
    // 是否信保亮灯Y/N
    IsCrdScrty   string `json:"is_crd_scrty,omitempty" xml:"is_crd_scrty,omitempty"`
    // 月PC在线时长达标率
    PcOnlineHourRate   string `json:"pc_online_hour_rate,omitempty" xml:"pc_online_hour_rate,omitempty"`
    // 近30天PC/MA在线时长(小时)
    StayHourPcM   string `json:"stay_hour_pc_m,omitempty" xml:"stay_hour_pc_m,omitempty"`
    // 月pc在线天数达标率
    PcOnlineDayRate   string `json:"pc_online_day_rate,omitempty" xml:"pc_online_day_rate,omitempty"`
    // 近30天PC/MA活跃天数
    ActiveDayPcM   int64 `json:"active_day_pc_m,omitempty" xml:"active_day_pc_m,omitempty"`
    // 月活跃子账号达标率
    ActiveMbrRate   string `json:"active_mbr_rate,omitempty" xml:"active_mbr_rate,omitempty"`
    // 近30天活跃子账号数
    ActiveMbrM   int64 `json:"active_mbr_m,omitempty" xml:"active_mbr_m,omitempty"`
    // 月mc达标率
    McRate   string `json:"mc_rate,omitempty" xml:"mc_rate,omitempty"`
    // 关键词达标率
    KwRate   string `json:"kw_rate,omitempty" xml:"kw_rate,omitempty"`
    // k200比率，关键词除以200的比率
    Kw200Rate   string `json:"kw200_rate,omitempty" xml:"kw200_rate,omitempty"`
    // 视频数量
    ProdVideoCnt   int64 `json:"prod_video_cnt,omitempty" xml:"prod_video_cnt,omitempty"`
    // 是否有侵权品牌Y/N
    IsTortBrand   string `json:"is_tort_brand,omitempty" xml:"is_tort_brand,omitempty"`
    // 是否有价格不合理产品Y/N
    HavePriceProd   string `json:"have_price_prod,omitempty" xml:"have_price_prod,omitempty"`
    // 是否有重复铺货产品Y/N
    HaveRepeatProd   string `json:"have_repeat_prod,omitempty" xml:"have_repeat_prod,omitempty"`
    // RTS商品数
    RtsProdCnt   int64 `json:"rts_prod_cnt,omitempty" xml:"rts_prod_cnt,omitempty"`
    // 实力优品达标率
    GoodProdRate   string `json:"good_prod_rate,omitempty" xml:"good_prod_rate,omitempty"`
    // 实力优品数
    GoodProdCnt   int64 `json:"good_prod_cnt,omitempty" xml:"good_prod_cnt,omitempty"`
    // 潜力品达标率
    PotentialProdRate   string `json:"potential_prod_rate,omitempty" xml:"potential_prod_rate,omitempty"`
    // 潜力品数
    PotentialProdCnt   int64 `json:"potential_prod_cnt,omitempty" xml:"potential_prod_cnt,omitempty"`
    // 蓝海品数
    ProdBlue   int64 `json:"prod_blue,omitempty" xml:"prod_blue,omitempty"`
    // 4.5分及以上商品数
    Prod45   int64 `json:"prod45,omitempty" xml:"prod45,omitempty"`
    // 4.0分及以上商品数
    Prod40   int64 `json:"prod40,omitempty" xml:"prod40,omitempty"`
    // p150且4分达标率
    P1504Rate   string `json:"p1504_rate,omitempty" xml:"p1504_rate,omitempty"`
    // p30且4分达标率
    P304Rate   string `json:"p304_rate,omitempty" xml:"p304_rate,omitempty"`
    // 有效商品数
    ProdCnt   int64 `json:"prod_cnt,omitempty" xml:"prod_cnt,omitempty"`
    // 店铺是否装修，页面最晚修改时间和旺铺创建时间差
    IsCompDec   string `json:"is_comp_dec,omitempty" xml:"is_comp_dec,omitempty"`
    // 服务商发品数
    AgencyProdCnt   int64 `json:"agency_prod_cnt,omitempty" xml:"agency_prod_cnt,omitempty"`
    // 订单完成时间
    TradesuccessTime   string `json:"tradesuccess_time,omitempty" xml:"tradesuccess_time,omitempty"`
    // 进入交付时间
    DeliveredTime   string `json:"delivered_time,omitempty" xml:"delivered_time,omitempty"`
    // 下单时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 交易单当前状态
    CurrentStatus   string `json:"current_status,omitempty" xml:"current_status,omitempty"`
    // 订单号
    OrdNum   string `json:"ord_num,omitempty" xml:"ord_num,omitempty"`
    // 服务名称
    ServiceName   string `json:"service_name,omitempty" xml:"service_name,omitempty"`
    // 服务编码
    ServiceCode   string `json:"service_code,omitempty" xml:"service_code,omitempty"`
    // 服务商编码
    VendorCode   string `json:"vendor_code,omitempty" xml:"vendor_code,omitempty"`
    // 服务商公司名
    VendorCompName   string `json:"vendor_comp_name,omitempty" xml:"vendor_comp_name,omitempty"`
    // 主营二级类目描述
    MainCateLv2Desc   string `json:"main_cate_lv2_desc,omitempty" xml:"main_cate_lv2_desc,omitempty"`
    // 主营二级类目ID
    MainCateLv2Std001   int64 `json:"main_cate_lv2_std001,omitempty" xml:"main_cate_lv2_std001,omitempty"`
    // 主营一级类目描述
    MainCateLv1Desc   string `json:"main_cate_lv1_desc,omitempty" xml:"main_cate_lv1_desc,omitempty"`
    // 主营一级类目ID
    MainCateLv1Std001   int64 `json:"main_cate_lv1_std001,omitempty" xml:"main_cate_lv1_std001,omitempty"`
    // 会员主账号id
    AdminMbrId   string `json:"admin_mbr_id,omitempty" xml:"admin_mbr_id,omitempty"`
    // 系统入库时间
    DwInsTime   string `json:"dw_ins_time,omitempty" xml:"dw_ins_time,omitempty"`
    // 统计日期
    StatDate   string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
    // RTS力是否双60
    Rts6060   string `json:"rts6060,omitempty" xml:"rts6060,omitempty"`
    // 当月星等级
    LevelStarPage   int64 `json:"level_star_page,omitempty" xml:"level_star_page,omitempty"`
    // 预测下月星等级
    LevelStar   int64 `json:"level_star,omitempty" xml:"level_star,omitempty"`
    // 最近7天店铺日均搜索曝光是否正常Y/N
    IsCompExpOk   string `json:"is_comp_exp_ok,omitempty" xml:"is_comp_exp_ok,omitempty"`
    // 关键词曝光是否正常Y/N
    IsKwExpOk   string `json:"is_kw_exp_ok,omitempty" xml:"is_kw_exp_ok,omitempty"`
    // 是否有降星风险Y/N
    IsStarLowerRisk   string `json:"is_star_lower_risk,omitempty" xml:"is_star_lower_risk,omitempty"`
    // 开通周期（付款到开通）
    OpenCycle   int64 `json:"open_cycle,omitempty" xml:"open_cycle,omitempty"`
    // 首年0-6个月信保渗透月数
    FiscalSvrc6mValidOrdCnt   int64 `json:"fiscal_svrc6m_valid_ord_cnt,omitempty" xml:"fiscal_svrc6m_valid_ord_cnt,omitempty"`
    // 是否p4p月消耗达标【Y/N】
    IsP4pCostQualified1   string `json:"is_p4p_cost_qualified1,omitempty" xml:"is_p4p_cost_qualified1,omitempty"`
    // 是否p4p绑定推广商品数大于等于80【Y/N】
    IsP4pProdGeq80Cnt   string `json:"is_p4p_prod_geq80_cnt,omitempty" xml:"is_p4p_prod_geq80_cnt,omitempty"`
    // 是否p4p绑定推广商品数大于等于30【Y/N】
    IsP4pProdGeq30Cnt   string `json:"is_p4p_prod_geq30_cnt,omitempty" xml:"is_p4p_prod_geq30_cnt,omitempty"`
    // 是否商机破蛋【Y/N】
    IsHasAb   string `json:"is_has_ab,omitempty" xml:"is_has_ab,omitempty"`
    // 最近一次开通服务天数
    LatestActualSrvcDays   int64 `json:"latest_actual_srvc_days,omitempty" xml:"latest_actual_srvc_days,omitempty"`
    // 是否首月询盘破蛋【Y/N】
    IsHasMcFstMon   string `json:"is_has_mc_fst_mon,omitempty" xml:"is_has_mc_fst_mon,omitempty"`
    // 是否MC15达标（询盘达到15个）
    IsMc15   string `json:"is_mc15,omitempty" xml:"is_mc15,omitempty"`
    // 服务商开始接洽时间
    ContactTime   string `json:"contact_time,omitempty" xml:"contact_time,omitempty"`
    // 是否蓝海定招
    IsBlueNewSign   string `json:"is_blue_new_sign,omitempty" xml:"is_blue_new_sign,omitempty"`
}
