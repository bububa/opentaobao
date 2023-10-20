package maitix

import (
	"sync"
)

// DisEncrypt4CmbResult 结构体
type DisEncrypt4CmbResult struct {
	// 订单金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 分行号
	BranchNo string `json:"branch_no,omitempty" xml:"branch_no,omitempty"`
	// 订单日期
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 请求日间
	DateTime string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	// 过期时间跨度
	ExpireTimeSpan string `json:"expire_time_span,omitempty" xml:"expire_time_span,omitempty"`
	// 商户号
	MerchantNo string `json:"merchant_no,omitempty" xml:"merchant_no,omitempty"`
	// 订单号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 支付成功通知的参数,这里传的是大麦的订单号
	PayNoticePara string `json:"pay_notice_para,omitempty" xml:"pay_notice_para,omitempty"`
	// 支付成功通知的地址
	PayNoticeUrl string `json:"pay_notice_url,omitempty" xml:"pay_notice_url,omitempty"`
	// 支付成功返回的地址
	ReturnUrl string `json:"return_url,omitempty" xml:"return_url,omitempty"`
	// 加密结果
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
}

var poolDisEncrypt4CmbResult = sync.Pool{
	New: func() any {
		return new(DisEncrypt4CmbResult)
	},
}

// GetDisEncrypt4CmbResult() 从对象池中获取DisEncrypt4CmbResult
func GetDisEncrypt4CmbResult() *DisEncrypt4CmbResult {
	return poolDisEncrypt4CmbResult.Get().(*DisEncrypt4CmbResult)
}

// ReleaseDisEncrypt4CmbResult 释放DisEncrypt4CmbResult
func ReleaseDisEncrypt4CmbResult(v *DisEncrypt4CmbResult) {
	v.Amount = ""
	v.BranchNo = ""
	v.Date = ""
	v.DateTime = ""
	v.ExpireTimeSpan = ""
	v.MerchantNo = ""
	v.OrderNo = ""
	v.PayNoticePara = ""
	v.PayNoticeUrl = ""
	v.ReturnUrl = ""
	v.Sign = ""
	poolDisEncrypt4CmbResult.Put(v)
}
