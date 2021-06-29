package koubeimall

// MallDTO 
type MallDTO struct {

    // 商圈LOGO
    
    MallLogo   string `json:"mall_logo,omitempty" xml:"mall_logo,omitempty"`
    

    // 商圈名称
    
    MallName   string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
    

    // 商圈联系方式，多个电话用英文","分隔符（包含手机、座机电话）
    
    ContactInfo   string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
    

    // 商圈封面
    
    MallCover   string `json:"mall_cover,omitempty" xml:"mall_cover,omitempty"`
    

    // 商圈ID
    
    MallId   string `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
    

    // 地理位置信息模型
    
    DistrictInfo   *DistrictInfo `json:"district_info,omitempty" xml:"district_info,omitempty"`
    

    // 商圈主页小程序URL，根据入参参数display_channel信息，获取对应端小程序URL，默认支付宝小程序链接
    
    MallHomePageUrl   string `json:"mall_home_page_url,omitempty" xml:"mall_home_page_url,omitempty"`
    

    // 营业时间
    
    BusinessTime   string `json:"business_time,omitempty" xml:"business_time,omitempty"`
    

    // 商圈类目，包含：购物中心、步行街、特色街区、火车站、百货、机场等
    
    CategoryName   string `json:"category_name,omitempty" xml:"category_name,omitempty"`
    

    // 商圈标签，多个标签用英文","分隔符
    
    MallLabel   string `json:"mall_label,omitempty" xml:"mall_label,omitempty"`
    

    // 距离
    
    Distance   string `json:"distance,omitempty" xml:"distance,omitempty"`
    

}
