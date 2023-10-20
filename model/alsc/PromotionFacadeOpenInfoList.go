package alsc

import (
	"sync"
)

// PromotionFacadeOpenInfoList 结构体
type PromotionFacadeOpenInfoList struct {
	// 圈选商品
	ItemSelectedOpenInfoList []ItemSelectedOpenInfoList `json:"item_selected_open_info_list,omitempty" xml:"item_selected_open_info_list>item_selected_open_info_list,omitempty"`
	// 圈选门店
	ShopSelectedOpenInfoList []ShopSelectedOpenInfoList `json:"shop_selected_open_info_list,omitempty" xml:"shop_selected_open_info_list>shop_selected_open_info_list,omitempty"`
	// 可用时段  0000101（星期五，星期天）       [{&#34;days&#34;:&#34;0000101&#34;,&#34;startTime&#34;:&#34;08:00:00&#34;,&#34;endTime&#34;:&#34;11:59:59&#34;},       {&#34;days&#34;:&#34;0010101&#34;,&#34;startTime&#34;:&#34;08:00:00&#34;,&#34;endTime&#34;:&#34;11:59:59&#34;}]
	AvailableTime string `json:"available_time,omitempty" xml:"available_time,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 规则描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 促销周期结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 适用商品范围  值：ALL，PART_AVAILABLE，PART_UNAVAILABLE      * 说明：全部商品可用，部分商品可用，部分商品不可用
	ItemCoverage string `json:"item_coverage,omitempty" xml:"item_coverage,omitempty"`
	// 促销规则名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 促销规则Id
	PromotionId string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// 促销周期开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 状态 促销活动状态 值：UNUSED,USED,NO_INVENTORY,INVALID 说明：未使用,使用中,使用中,使用中
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 促销活动适用人群 值：MEMBER，CUSTOMER，ALL      * 说明：会员,非会员，不限
	SuitablePeople string `json:"suitable_people,omitempty" xml:"suitable_people,omitempty"`
	// 更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// *      * 满量促销           TYPE_FULL_AMOUNT,      *      * 满额促销           TYPE_FULL_CAPACITY,      *      * 买赠活动           TYPE_BOUGHT_GIFT;
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 扩展字段  isAdded:是否叠加,isVoucherShared:是否与优惠券共享,giftGoodsIdList:赠品类似、活动商品,privilegeCondition:权益条件(类型type:         满量 &#34;type&#34;:&#34;FULL_AMOUNT&#34;,&#34;name&#34;:&#34;num&#34;, &#34;value&#34;:&#34;3&#34;         满额 &#34;type&#34;:&#34;FULL_CAPACITY&#34;, &#34;name&#34;:&#34;money&#34;,&#34;value&#34;:&#34;300&#34;         下一份 &#34;type&#34;:&#34;NEXT&#34;, &#34;name&#34;:&#34;num&#34;,&#34;value&#34;:&#34;3&#34;         加价购),privilegeType:权益类型(一口价&#34;type&#34;:&#34;FIXPRICE&#34;,&#34;name&#34;:&#34;money&#34;,&#34;value&#34;:&#34;3000&#34;         减免&#34;type&#34;:&#34;DECREASE&#34;,&#34;name&#34;:&#34;money&#34;,&#34;value&#34;:&#34;10&#34;         减低价&#34;type&#34;:&#34;REDUCE_LOW_PRICE&#34;,&#34;name&#34;:&#34;&#34;,&#34;value&#34;:&#34;&#34;         折扣：&#34;type&#34;:&#34;DISCOUNT&#34;,&#34;name&#34;:&#34;discount&#34;,&#34;value&#34;:&#34;80&#34;         赠品 &#34;type&#34;:&#34;GIFT&#34;, &#34;name&#34;:&#34;num&#34;, &#34;value&#34;:&#34;2&#34;)
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 创建人
	CreateBy string `json:"create_by,omitempty" xml:"create_by,omitempty"`
	// 更新人
	UpdateBy string `json:"update_by,omitempty" xml:"update_by,omitempty"`
	// 更新人名称
	UpdateByName string `json:"update_by_name,omitempty" xml:"update_by_name,omitempty"`
	// 是否已经删除
	Deleted bool `json:"deleted,omitempty" xml:"deleted,omitempty"`
}

var poolPromotionFacadeOpenInfoList = sync.Pool{
	New: func() any {
		return new(PromotionFacadeOpenInfoList)
	},
}

// GetPromotionFacadeOpenInfoList() 从对象池中获取PromotionFacadeOpenInfoList
func GetPromotionFacadeOpenInfoList() *PromotionFacadeOpenInfoList {
	return poolPromotionFacadeOpenInfoList.Get().(*PromotionFacadeOpenInfoList)
}

// ReleasePromotionFacadeOpenInfoList 释放PromotionFacadeOpenInfoList
func ReleasePromotionFacadeOpenInfoList(v *PromotionFacadeOpenInfoList) {
	v.ItemSelectedOpenInfoList = v.ItemSelectedOpenInfoList[:0]
	v.ShopSelectedOpenInfoList = v.ShopSelectedOpenInfoList[:0]
	v.AvailableTime = ""
	v.GmtCreate = ""
	v.Description = ""
	v.EndTime = ""
	v.ItemCoverage = ""
	v.Name = ""
	v.PromotionId = ""
	v.StartTime = ""
	v.Status = ""
	v.SuitablePeople = ""
	v.GmtModified = ""
	v.Type = ""
	v.ExtInfo = ""
	v.CreateBy = ""
	v.UpdateBy = ""
	v.UpdateByName = ""
	v.Deleted = false
	poolPromotionFacadeOpenInfoList.Put(v)
}
