package alitripmerchant

// ListSearchParam 
type ListSearchParam struct {
    // 最高价格
    PriceMax   int64 `json:"price_max,omitempty" xml:"price_max,omitempty"`
    // 成人数
    AdultNum   int64 `json:"adult_num,omitempty" xml:"adult_num,omitempty"`
    // 儿童数
    ChildNum   int64 `json:"child_num,omitempty" xml:"child_num,omitempty"`
    // 偏移量
    Offset   int64 `json:"offset,omitempty" xml:"offset,omitempty"`
    // 星级筛选
    Star   string `json:"star,omitempty" xml:"star,omitempty"`
    // 城市编码，默认上海
    CityCode   string `json:"city_code,omitempty" xml:"city_code,omitempty"`
    // 会员等级
    MemberLevel   string `json:"member_level,omitempty" xml:"member_level,omitempty"`
    // 每页数量
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 升降序 0降序 1升序
    Dir   int64 `json:"dir,omitempty" xml:"dir,omitempty"`
    // 儿童年龄
    ChildrenAges   string `json:"children_ages,omitempty" xml:"children_ages,omitempty"`
    // 最低价格
    PriceMin   int64 `json:"price_min,omitempty" xml:"price_min,omitempty"`
    // 入店时间
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    // 当前页
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // 离店时间
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    // 品牌编码
    Brand   string `json:"brand,omitempty" xml:"brand,omitempty"`
    // 用户登录信息
    Token   string `json:"token,omitempty" xml:"token,omitempty"`
}
