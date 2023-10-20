package normalvisa

import (
	"sync"
)

// NormalVisaInfo 结构体
type NormalVisaInfo struct {
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 单价
	AuctionPrice string `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 国家id，国家编码详见：https://open.alitrip.com/docs/doc.htm?spm=a21tt.7629140.0.0.retXmq&amp;treeId=79&amp;articleId=104840&amp;docType=1
	CountryId string `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 买家ID
	Openuid string `json:"openuid,omitempty" xml:"openuid,omitempty"`
	// 1:贴纸签  2:电子签  3:面试
	NormalVisaType int64 `json:"normal_visa_type,omitempty" xml:"normal_visa_type,omitempty"`
	// 无办理人信息：1001 办理人已填写：1002 已收到资料: 1003 已审核完成: 1004 已送签:1005 结果已返回: 1006 已预约面试: 1007 处理中:1008 已中止办理：1010
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 是否达到中止状态
	EndStatus bool `json:"end_status,omitempty" xml:"end_status,omitempty"`
	// 是否需要商家代填
	NeedFillContact bool `json:"need_fill_contact,omitempty" xml:"need_fill_contact,omitempty"`
}

var poolNormalVisaInfo = sync.Pool{
	New: func() any {
		return new(NormalVisaInfo)
	},
}

// GetNormalVisaInfo() 从对象池中获取NormalVisaInfo
func GetNormalVisaInfo() *NormalVisaInfo {
	return poolNormalVisaInfo.Get().(*NormalVisaInfo)
}

// ReleaseNormalVisaInfo 释放NormalVisaInfo
func ReleaseNormalVisaInfo(v *NormalVisaInfo) {
	v.PayTime = ""
	v.AuctionPrice = ""
	v.Title = ""
	v.StatusDesc = ""
	v.CountryId = ""
	v.Openuid = ""
	v.NormalVisaType = 0
	v.Status = 0
	v.Amount = 0
	v.BizOrderId = 0
	v.EndStatus = false
	v.NeedFillContact = false
	poolNormalVisaInfo.Put(v)
}
