package scs

import (
	"sync"
)

// ReportResultTopDto 结构体
type ReportResultTopDto struct {
	// 按天维度划分的日期
	LogDate string `json:"log_date,omitempty" xml:"log_date,omitempty"`
	// 点击率
	Ctr string `json:"ctr,omitempty" xml:"ctr,omitempty"`
	// 千次展现成本
	Ecpm string `json:"ecpm,omitempty" xml:"ecpm,omitempty"`
	// 消耗
	Charge string `json:"charge,omitempty" xml:"charge,omitempty"`
	// 点击成本
	Ecpc string `json:"ecpc,omitempty" xml:"ecpc,omitempty"`
	// 成交金额
	AlipayInshopAmt string `json:"alipay_inshop_amt,omitempty" xml:"alipay_inshop_amt,omitempty"`
	// 成交转化率
	Cvr string `json:"cvr,omitempty" xml:"cvr,omitempty"`
	// ROI
	Roi string `json:"roi,omitempty" xml:"roi,omitempty"`
	// 预售成交笔数
	PrepayInshopAmt string `json:"prepay_inshop_amt,omitempty" xml:"prepay_inshop_amt,omitempty"`
	// 新客成交金额贡献占比
	NoLalipayInshopAmtProprtion string `json:"no_lalipay_inshop_amt_proprtion,omitempty" xml:"no_lalipay_inshop_amt_proprtion,omitempty"`
	// 直接成交金额
	DirAlipayInshopAmt string `json:"dir_alipay_inshop_amt,omitempty" xml:"dir_alipay_inshop_amt,omitempty"`
	// 间接成交金额
	IndirAlipayInshopAmt string `json:"indir_alipay_inshop_amt,omitempty" xml:"indir_alipay_inshop_amt,omitempty"`
	// 派样成交金额
	SampleAlipayAmt string `json:"sample_alipay_amt,omitempty" xml:"sample_alipay_amt,omitempty"`
	// 成交金额
	AlipayInshopAmtKuan string `json:"alipay_inshop_amt_kuan,omitempty" xml:"alipay_inshop_amt_kuan,omitempty"`
	// 成交转化率
	CvrKuan string `json:"cvr_kuan,omitempty" xml:"cvr_kuan,omitempty"`
	// ROI
	RoiKuan string `json:"roi_kuan,omitempty" xml:"roi_kuan,omitempty"`
	// 预售成交笔数
	PrepayInshopAmtKuan string `json:"prepay_inshop_amt_kuan,omitempty" xml:"prepay_inshop_amt_kuan,omitempty"`
	// 新客成交金额贡献占比
	NoLalipayInshopAmtProprtionKuan string `json:"no_lalipay_inshop_amt_proprtion_kuan,omitempty" xml:"no_lalipay_inshop_amt_proprtion_kuan,omitempty"`
	// 直接成交金额
	DirAlipayInshopAmtKuan string `json:"dir_alipay_inshop_amt_kuan,omitempty" xml:"dir_alipay_inshop_amt_kuan,omitempty"`
	// 间接成交金额
	IndirAlipayInshopAmtKuan string `json:"indir_alipay_inshop_amt_kuan,omitempty" xml:"indir_alipay_inshop_amt_kuan,omitempty"`
	// 派样成交金额
	SampleAlipayAmtKuan string `json:"sample_alipay_amt_kuan,omitempty" xml:"sample_alipay_amt_kuan,omitempty"`
	// 按小时维度划分的小时ID
	HourId int64 `json:"hour_id,omitempty" xml:"hour_id,omitempty"`
	// 展现数
	AdPv int64 `json:"ad_pv,omitempty" xml:"ad_pv,omitempty"`
	// 点击数
	Click int64 `json:"click,omitempty" xml:"click,omitempty"`
	// 加购量
	CarNum int64 `json:"car_num,omitempty" xml:"car_num,omitempty"`
	// 直接加购量
	DirCarNum int64 `json:"dir_car_num,omitempty" xml:"dir_car_num,omitempty"`
	// 间接加购量
	IndirCarNum int64 `json:"indir_car_num,omitempty" xml:"indir_car_num,omitempty"`
	// 收藏量
	InshopItemColNum int64 `json:"inshop_item_col_num,omitempty" xml:"inshop_item_col_num,omitempty"`
	// 收藏加购成本
	InshopItemColCarNumCost int64 `json:"inshop_item_col_car_num_cost,omitempty" xml:"inshop_item_col_car_num_cost,omitempty"`
	// 成交笔数
	AlipayInshopNum int64 `json:"alipay_inshop_num,omitempty" xml:"alipay_inshop_num,omitempty"`
	// 预售成交金额
	PrepayInshopNum int64 `json:"prepay_inshop_num,omitempty" xml:"prepay_inshop_num,omitempty"`
	// 直接成交笔数
	DirAlipayInshopNum int64 `json:"dir_alipay_inshop_num,omitempty" xml:"dir_alipay_inshop_num,omitempty"`
	// 间接成交笔数
	IndirAlipayInshopNum int64 `json:"indir_alipay_inshop_num,omitempty" xml:"indir_alipay_inshop_num,omitempty"`
	// 派样量
	SampleAlipayNum int64 `json:"sample_alipay_num,omitempty" xml:"sample_alipay_num,omitempty"`
	// 加购量
	CarNumKuan int64 `json:"car_num_kuan,omitempty" xml:"car_num_kuan,omitempty"`
	// 直接加购量
	DirCarNumKuan int64 `json:"dir_car_num_kuan,omitempty" xml:"dir_car_num_kuan,omitempty"`
	// 间接加购量
	IndirCarNumKuan int64 `json:"indir_car_num_kuan,omitempty" xml:"indir_car_num_kuan,omitempty"`
	// 收藏量
	InshopItemColNumKuan int64 `json:"inshop_item_col_num_kuan,omitempty" xml:"inshop_item_col_num_kuan,omitempty"`
	// 收藏加购成本
	InshopItemColCarNumCostKuan int64 `json:"inshop_item_col_car_num_cost_kuan,omitempty" xml:"inshop_item_col_car_num_cost_kuan,omitempty"`
	// 成交笔数
	AlipayInshopNumKuan int64 `json:"alipay_inshop_num_kuan,omitempty" xml:"alipay_inshop_num_kuan,omitempty"`
	// 预售成交金额
	PrepayInshopNumKuan int64 `json:"prepay_inshop_num_kuan,omitempty" xml:"prepay_inshop_num_kuan,omitempty"`
	// 直接成交笔数
	DirAlipayInshopNumKuan int64 `json:"dir_alipay_inshop_num_kuan,omitempty" xml:"dir_alipay_inshop_num_kuan,omitempty"`
	// 间接成交笔数
	IndirAlipayInshopNumKuan int64 `json:"indir_alipay_inshop_num_kuan,omitempty" xml:"indir_alipay_inshop_num_kuan,omitempty"`
	// 派样量
	SampleAlipayNumKuan int64 `json:"sample_alipay_num_kuan,omitempty" xml:"sample_alipay_num_kuan,omitempty"`
	// 计划ID
	CampaignId int64 `json:"campaign_id,omitempty" xml:"campaign_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolReportResultTopDto = sync.Pool{
	New: func() any {
		return new(ReportResultTopDto)
	},
}

// GetReportResultTopDto() 从对象池中获取ReportResultTopDto
func GetReportResultTopDto() *ReportResultTopDto {
	return poolReportResultTopDto.Get().(*ReportResultTopDto)
}

// ReleaseReportResultTopDto 释放ReportResultTopDto
func ReleaseReportResultTopDto(v *ReportResultTopDto) {
	v.LogDate = ""
	v.Ctr = ""
	v.Ecpm = ""
	v.Charge = ""
	v.Ecpc = ""
	v.AlipayInshopAmt = ""
	v.Cvr = ""
	v.Roi = ""
	v.PrepayInshopAmt = ""
	v.NoLalipayInshopAmtProprtion = ""
	v.DirAlipayInshopAmt = ""
	v.IndirAlipayInshopAmt = ""
	v.SampleAlipayAmt = ""
	v.AlipayInshopAmtKuan = ""
	v.CvrKuan = ""
	v.RoiKuan = ""
	v.PrepayInshopAmtKuan = ""
	v.NoLalipayInshopAmtProprtionKuan = ""
	v.DirAlipayInshopAmtKuan = ""
	v.IndirAlipayInshopAmtKuan = ""
	v.SampleAlipayAmtKuan = ""
	v.HourId = 0
	v.AdPv = 0
	v.Click = 0
	v.CarNum = 0
	v.DirCarNum = 0
	v.IndirCarNum = 0
	v.InshopItemColNum = 0
	v.InshopItemColCarNumCost = 0
	v.AlipayInshopNum = 0
	v.PrepayInshopNum = 0
	v.DirAlipayInshopNum = 0
	v.IndirAlipayInshopNum = 0
	v.SampleAlipayNum = 0
	v.CarNumKuan = 0
	v.DirCarNumKuan = 0
	v.IndirCarNumKuan = 0
	v.InshopItemColNumKuan = 0
	v.InshopItemColCarNumCostKuan = 0
	v.AlipayInshopNumKuan = 0
	v.PrepayInshopNumKuan = 0
	v.DirAlipayInshopNumKuan = 0
	v.IndirAlipayInshopNumKuan = 0
	v.SampleAlipayNumKuan = 0
	v.CampaignId = 0
	v.ItemId = 0
	poolReportResultTopDto.Put(v)
}
