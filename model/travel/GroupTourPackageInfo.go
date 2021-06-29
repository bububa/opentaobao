package travel

// GroupTourPackageInfo 
type GroupTourPackageInfo struct {
    // 行程信息，新发布商品必填，数据，可以一次提交多天的行程信息
    TripElementList   []GroupTourTripElement `json:"trip_element_list,omitempty" xml:"trip_element_list>group_tour_trip_element,omitempty"`
    // 返程交通，交通工具的类型必须与商品上的去程交通一致
    BackTrafficInfoList   []GroupTourTrafficInfo `json:"back_traffic_info_list,omitempty" xml:"back_traffic_info_list>group_tour_traffic_info,omitempty"`
    // 去程交通，交通工具的类型必须与商品上的去程交通一致
    GoTrafficInfoList   []GroupTourTrafficInfo `json:"go_traffic_info_list,omitempty" xml:"go_traffic_info_list>group_tour_traffic_info,omitempty"`
    // 新发布商品时必填。自费项目。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    SelfExplanation   []string `json:"self_explanation,omitempty" xml:"self_explanation>string,omitempty"`
    // 新发布商品时必填。买家预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    OrderInfo   []string `json:"order_info,omitempty" xml:"order_info>string,omitempty"`
    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    FeeExclude   []string `json:"fee_exclude,omitempty" xml:"fee_exclude>string,omitempty"`
    // 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    FeeInclude   []string `json:"fee_include,omitempty" xml:"fee_include>string,omitempty"`
    // 套餐对应的出发地，是商品出发地的子集
    FromLocations   string `json:"from_locations,omitempty" xml:"from_locations,omitempty"`
    // 必填，跟团游套餐对应的商家编码
    OutProductId   string `json:"out_product_id,omitempty" xml:"out_product_id,omitempty"`
    // 新发布商品必填，跟团游套餐名
    PackageName   string `json:"package_name,omitempty" xml:"package_name,omitempty"`
    // 套餐操作类型(0:套餐覆盖修改,1:增加套餐,2:删除套餐)===默认为0===
    PackageOperation   int64 `json:"package_operation,omitempty" xml:"package_operation,omitempty"`
    // 套餐级别行程天数，必填。
    TripDays   int64 `json:"trip_days,omitempty" xml:"trip_days,omitempty"`
    // 套餐级别行程晚数，必填。
    AccomNights   int64 `json:"accom_nights,omitempty" xml:"accom_nights,omitempty"`
}
