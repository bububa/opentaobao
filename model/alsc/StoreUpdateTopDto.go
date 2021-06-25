package alsc

// StoreUpdateTopDto 
type StoreUpdateTopDto struct {

    // 门店主名
    Name   string `json:"name,omitempty"`

    // 分店名称
    SubName   string `json:"sub_name,omitempty"`

    // 门店主类目
    MainCategory   int64 `json:"main_category,omitempty"`

    // 门店外部编码
    OuterCode   string `json:"outer_code,omitempty"`

    // 门店类型，枚举值。NORMAL：普通门店。暂时统一使用这个值
    StoreType   string `json:"store_type,omitempty"`

    // 门店开始营业时间
    StartTime   string `json:"start_time,omitempty"`

    // 门店结束营业时间
    EndTime   string `json:"end_time,omitempty"`

    // 门店地址
    StoreAdress   *StoreAdressDto `json:"store_adress,omitempty"`

    // 门店信息
    StoreKeeper   *StoreKeeperDto `json:"store_keeper,omitempty"`

    // 门店状态，枚举值。NORMAL：正常。CLOSE：关店。HOLD: 暂停营业
    Status   string `json:"status,omitempty"`

    // 备注
    Description   string `json:"description,omitempty"`

    // 店铺id
    ShopId   int64 `json:"shop_id,omitempty"`

    // 门店类型
    BizType   int64 `json:"biz_type,omitempty"`

    // 门店标
    Tags   []Number `json:"tags,omitempty"`

    // 24小时营业
    AllDay   bool `json:"all_day,omitempty"`

    // 通用属性
    Attributes   []AttributeValueTopDto `json:"attributes,omitempty"`

    // 类目属性
    CategoryPropertyValues   []PropertyValueTopDto `json:"category_property_values,omitempty"`

    // 标准类目ID
    StandardCategoryId   string `json:"standard_category_id,omitempty"`

    // 业务身份
    BizCode   string `json:"biz_code,omitempty"`

    // 业务属性
    BizAttributes   []AttributeValueTopDto `json:"biz_attributes,omitempty"`

    // 门店logo
    Logo   string `json:"logo,omitempty"`

    // 是否v3
    IsV3   bool `json:"is_v3,omitempty"`

    // 星期
    Week   []String `json:"week,omitempty"`

    // 门店头图
    Pic   string `json:"pic,omitempty"`

    // 审核状态
    AuthenStatus   int64 `json:"authen_status,omitempty"`

    // 门店id
    StoreId   int64 `json:"store_id,omitempty"`

}
