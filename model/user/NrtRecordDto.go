package user

// NrtRecordDto 
type NrtRecordDto struct {

    // 业务类型
    BizCode   string `json:"biz_code,omitempty"`

    // 省编码
    ProvinceCode   string `json:"province_code,omitempty"`

    // 市编码
    CityCode   string `json:"city_code,omitempty"`

    // 0-预约到店  4-折扣卡  5-有价礼包  11-一口价订单  14-POS订单  17-其他
    Channel   int64 `json:"channel,omitempty"`

    // 创建人ID
    CreatorId   int64 `json:"creator_id,omitempty"`

    // 创建人姓名
    CreatorName   string `json:"creator_name,omitempty"`

    // 导购员ID
    EmployeeId   int64 `json:"employee_id,omitempty"`

    // 淘宝ID，取不到时可空
    BuyerId   string `json:"buyer_id,omitempty"`

    // 店铺ID
    StoreId   int64 `json:"store_id,omitempty"`

    // 区编码
    AreaCode   string `json:"area_code,omitempty"`

    // 市名
    CityName   string `json:"city_name,omitempty"`

    // 客户手机号
    Phone   string `json:"phone,omitempty"`

    // 区名
    AreaName   string `json:"area_name,omitempty"`

    // 客户姓名
    Name   string `json:"name,omitempty"`

    // 外部客资记录唯一ID
    OuterId   string `json:"outer_id,omitempty"`

    // 省名
    ProvinceName   string `json:"province_name,omitempty"`

    // 客户地址
    Address   string `json:"address,omitempty"`

}
