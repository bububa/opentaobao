package alitripmerchant

import (
	"sync"
)

// OfferDetailsDto 结构体
type OfferDetailsDto struct {
	// 酒店数据集合
	HotelContentList []HotelContentDto `json:"hotel_content_list,omitempty" xml:"hotel_content_list>hotel_content_dto,omitempty"`
	// 关联品牌
	JoinBrands []JoinBrandDto `json:"join_brands,omitempty" xml:"join_brands>join_brand_dto,omitempty"`
	// 无线端图片
	OfferImageWireless []string `json:"offer_image_wireless,omitempty" xml:"offer_image_wireless>string,omitempty"`
	// offer名称
	OfferName string `json:"offer_name,omitempty" xml:"offer_name,omitempty"`
	// 子标题
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 预定开始时间
	BookStartDate string `json:"book_start_date,omitempty" xml:"book_start_date,omitempty"`
	// 预定结束时间
	BookEndDate string `json:"book_end_date,omitempty" xml:"book_end_date,omitempty"`
	// 起价单位
	FromPriceUnit string `json:"from_price_unit,omitempty" xml:"from_price_unit,omitempty"`
	// 小图
	OfferImageSmall string `json:"offer_image_small,omitempty" xml:"offer_image_small,omitempty"`
	// 大图
	OfferImageBig string `json:"offer_image_big,omitempty" xml:"offer_image_big,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// offer跳转url
	OfferRedirectUrl string `json:"offer_redirect_url,omitempty" xml:"offer_redirect_url,omitempty"`
	// 起价
	FromPriceAmount string `json:"from_price_amount,omitempty" xml:"from_price_amount,omitempty"`
	// 时区
	TimeZone string `json:"time_zone,omitempty" xml:"time_zone,omitempty"`
	// offer类型
	OfferType string `json:"offer_type,omitempty" xml:"offer_type,omitempty"`
	// 入住开始时间
	ClEndDate string `json:"cl_end_date,omitempty" xml:"cl_end_date,omitempty"`
	// 入住结束时间
	ClStartDate string `json:"cl_start_date,omitempty" xml:"cl_start_date,omitempty"`
	// 子类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 活动底部点击按钮文案
	ButtonContent string `json:"button_content,omitempty" xml:"button_content,omitempty"`
	// offerId
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 抽奖活动
	ActivityLuckyDraw *ActivityLuckyDrawVo `json:"activity_lucky_draw,omitempty" xml:"activity_lucky_draw,omitempty"`
	// 弹屏信息
	PopUpInfo *PopUpInfo `json:"pop_up_info,omitempty" xml:"pop_up_info,omitempty"`
	// 是否配置了活动详情页
	NeedActivityPage bool `json:"need_activity_page,omitempty" xml:"need_activity_page,omitempty"`
	// 是否配置了弹屏
	NeedPopUpScreen bool `json:"need_pop_up_screen,omitempty" xml:"need_pop_up_screen,omitempty"`
}

var poolOfferDetailsDto = sync.Pool{
	New: func() any {
		return new(OfferDetailsDto)
	},
}

// GetOfferDetailsDto() 从对象池中获取OfferDetailsDto
func GetOfferDetailsDto() *OfferDetailsDto {
	return poolOfferDetailsDto.Get().(*OfferDetailsDto)
}

// ReleaseOfferDetailsDto 释放OfferDetailsDto
func ReleaseOfferDetailsDto(v *OfferDetailsDto) {
	v.HotelContentList = v.HotelContentList[:0]
	v.JoinBrands = v.JoinBrands[:0]
	v.OfferImageWireless = v.OfferImageWireless[:0]
	v.OfferName = ""
	v.Subtitle = ""
	v.BookStartDate = ""
	v.BookEndDate = ""
	v.FromPriceUnit = ""
	v.OfferImageSmall = ""
	v.OfferImageBig = ""
	v.Description = ""
	v.OfferRedirectUrl = ""
	v.FromPriceAmount = ""
	v.TimeZone = ""
	v.OfferType = ""
	v.ClEndDate = ""
	v.ClStartDate = ""
	v.SubType = ""
	v.ButtonContent = ""
	v.OfferId = 0
	v.ActivityLuckyDraw = nil
	v.PopUpInfo = nil
	v.NeedActivityPage = false
	v.NeedPopUpScreen = false
	poolOfferDetailsDto.Put(v)
}
