package alitripmerchant

// OfferDetailsDTO 
type OfferDetailsDTO struct {
    // offerId
    OfferId   int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
    // offer名称
    OfferName   string `json:"offer_name,omitempty" xml:"offer_name,omitempty"`
    // 子标题
    Subtitle   string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
    // 起价单位
    FromPriceUnit   string `json:"from_price_unit,omitempty" xml:"from_price_unit,omitempty"`
    // 小图
    OfferImageSmall   string `json:"offer_image_small,omitempty" xml:"offer_image_small,omitempty"`
    // 大图
    OfferImageBig   string `json:"offer_image_big,omitempty" xml:"offer_image_big,omitempty"`
    // 描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // offer跳转url
    OfferRedirectUrl   string `json:"offer_redirect_url,omitempty" xml:"offer_redirect_url,omitempty"`
    // 关联品牌
    JoinBrands   []JoinBrandDTO `json:"join_brands,omitempty" xml:"join_brands>join_brand_dto,omitempty"`
    // 起价
    FromPriceAmount   string `json:"from_price_amount,omitempty" xml:"from_price_amount,omitempty"`
    // 时区
    TimeZone   string `json:"time_zone,omitempty" xml:"time_zone,omitempty"`
    // offer类型
    OfferType   string `json:"offer_type,omitempty" xml:"offer_type,omitempty"`
    // 预定开始时间
    BookStartDate   string `json:"book_start_date,omitempty" xml:"book_start_date,omitempty"`
    // 预定结束时间
    BookEndDate   string `json:"book_end_date,omitempty" xml:"book_end_date,omitempty"`
    // 入住开始时间
    ClEndDate   string `json:"cl_end_date,omitempty" xml:"cl_end_date,omitempty"`
    // 入住结束时间
    ClStartDate   string `json:"cl_start_date,omitempty" xml:"cl_start_date,omitempty"`
    // 子类型
    SubType   string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
    // 无线端图片
    OfferImageWireless   []string `json:"offer_image_wireless,omitempty" xml:"offer_image_wireless>string,omitempty"`
    // 酒店数据集合
    HotelContentList   []HotelContentDTO `json:"hotel_content_list,omitempty" xml:"hotel_content_list>hotel_content_dto,omitempty"`
}
