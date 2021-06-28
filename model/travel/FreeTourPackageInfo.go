package travel

// FreeTourPackageInfo 
/* model for simplify = false
type FreeTourPackageInfo struct {

    // 必须传，套餐对应的商家编码
    
    OutProductId   string `json:"out_product_id,omitempty"`
    

    // 包含元素-景区门票，如果该套餐包含景区门票，则需要传这个参数
    
    FreeTourScenicInfoList  struct {
        FreeTourScenicInfo  []FreeTourScenicInfo `json:"free_tour_scenic_info,omitempty"`
    } `json:"free_tour_scenic_info_list,omitempty"`
    

    // 必填，套餐名称
    
    PackageName   string `json:"package_name,omitempty"`
    

    // 发布商品时必填。套餐对应的目的地，套餐对应的目的地必须是商品的目的地的子集，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称 ）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性 批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124>）会根据映射关系自动转换成飞猪标准地址
    
    ToLocations   string `json:"to_locations,omitempty"`
    

    // 返程交通，交通工具必须与商品上的返程交通一致
    
    BackTrafficInfoList  struct {
        FreeTourTrafficInfo  []FreeTourTrafficInfo `json:"free_tour_traffic_info,omitempty"`
    } `json:"back_traffic_info_list,omitempty"`
    

    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    
    FeeExclude  struct {
        String  []string `json:"string,omitempty"`
    } `json:"fee_exclude,omitempty"`
    

    // 门票说明
    
    ScenicDesc   string `json:"scenic_desc,omitempty"`
    

    // 去程交通，交通工具的类型必须与商品上的去程交通一致
    
    GoTrafficInfoList  struct {
        FreeTourTrafficInfo  []FreeTourTrafficInfo `json:"free_tour_traffic_info,omitempty"`
    } `json:"go_traffic_info_list,omitempty"`
    

    // 买家预定须知
    
    OrderInfo  struct {
        String  []string `json:"string,omitempty"`
    } `json:"order_info,omitempty"`
    

    // 当包含元素选择了餐饮，自驾租车，保险服务，其他费用时，该项必填
    
    ItemResourceInfoList  struct {
        ItemResourceInfo  []ItemResourceInfo `json:"item_resource_info,omitempty"`
    } `json:"item_resource_info_list,omitempty"`
    

    // 套餐对应的出发地，是商品出发地的子集
    
    FromLocations   string `json:"from_locations,omitempty"`
    

    // 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    
    FeeInclude  struct {
        String  []string `json:"string,omitempty"`
    } `json:"fee_include,omitempty"`
    

    // 当包含元素选择了酒店的时候，该项必填
    
    FreeTourHotelInfoList  struct {
        FreeTourHotelInfo  []FreeTourHotelInfo `json:"free_tour_hotel_info,omitempty"`
    } `json:"free_tour_hotel_info_list,omitempty"`
    

    // 套餐操作类型，(0:套餐覆盖修改,1:增加套餐,2:删除套餐)===默认为0===
    
    PackageOperation   int64 `json:"package_operation,omitempty"`
    

}
*/

// FreeTourPackageInfo 
type FreeTourPackageInfo struct {

    // 必须传，套餐对应的商家编码
    OutProductId   string `json:"out_product_id,omitempty"`

    // 包含元素-景区门票，如果该套餐包含景区门票，则需要传这个参数
    FreeTourScenicInfoList   []FreeTourScenicInfo `json:"free_tour_scenic_info_list,omitempty"`

    // 必填，套餐名称
    PackageName   string `json:"package_name,omitempty"`

    // 发布商品时必填。套餐对应的目的地，套餐对应的目的地必须是商品的目的地的子集，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称 ）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性 批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124>）会根据映射关系自动转换成飞猪标准地址
    ToLocations   string `json:"to_locations,omitempty"`

    // 返程交通，交通工具必须与商品上的返程交通一致
    BackTrafficInfoList   []FreeTourTrafficInfo `json:"back_traffic_info_list,omitempty"`

    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    FeeExclude   []string `json:"fee_exclude,omitempty"`

    // 门票说明
    ScenicDesc   string `json:"scenic_desc,omitempty"`

    // 去程交通，交通工具的类型必须与商品上的去程交通一致
    GoTrafficInfoList   []FreeTourTrafficInfo `json:"go_traffic_info_list,omitempty"`

    // 买家预定须知
    OrderInfo   []string `json:"order_info,omitempty"`

    // 当包含元素选择了餐饮，自驾租车，保险服务，其他费用时，该项必填
    ItemResourceInfoList   []ItemResourceInfo `json:"item_resource_info_list,omitempty"`

    // 套餐对应的出发地，是商品出发地的子集
    FromLocations   string `json:"from_locations,omitempty"`

    // 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    FeeInclude   []string `json:"fee_include,omitempty"`

    // 当包含元素选择了酒店的时候，该项必填
    FreeTourHotelInfoList   []FreeTourHotelInfo `json:"free_tour_hotel_info_list,omitempty"`

    // 套餐操作类型，(0:套餐覆盖修改,1:增加套餐,2:删除套餐)===默认为0===
    PackageOperation   int64 `json:"package_operation,omitempty"`

}
