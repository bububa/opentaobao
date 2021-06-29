package koubeimall

// StoreDto 
type StoreDto struct {
    // 门店品牌名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 门店联系方式，多个电话用英文","分隔符（包含手机、座机电话）
    ContactInfo   string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
    // 门店服务 预定、点餐、排号、外卖
    ServiceTagList   []ServiceTag `json:"service_tag_list,omitempty" xml:"service_tag_list>service_tag,omitempty"`
    // 门店所属商圈ID
    MallId   string `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
    // 门店评论总数
    CommentTotalCount   string `json:"comment_total_count,omitempty" xml:"comment_total_count,omitempty"`
    // 地理位置信息模型
    DistrictInfo   *DistrictInfo `json:"district_info,omitempty" xml:"district_info,omitempty"`
    // 门店评分
    CommentScore   string `json:"comment_score,omitempty" xml:"comment_score,omitempty"`
    // 营业时间
    BusinessTime   string `json:"business_time,omitempty" xml:"business_time,omitempty"`
    // 门店ID
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 榜单
    Billboard   string `json:"billboard,omitempty" xml:"billboard,omitempty"`
    // 门店类目信息，一级、二级、三级类目用英文","分隔符；示例：美食,火锅,川菜/重庆火锅
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    // 门店质量分
    Score   string `json:"score,omitempty" xml:"score,omitempty"`
    // 门店详情小程序URL
    StoreDetailUrl   string `json:"store_detail_url,omitempty" xml:"store_detail_url,omitempty"`
    // 门店logo
    StoreLogo   string `json:"store_logo,omitempty" xml:"store_logo,omitempty"`
    // 门店名称
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`
    // 人均消费
    AveragePrice   string `json:"average_price,omitempty" xml:"average_price,omitempty"`
}
