package alsc

// VoucherTemplateOpenInfo 
/* model for simplify = false
type VoucherTemplateOpenInfo struct {

    // {\"type\":0,\"settings\":[{\"dayStartTime\":\"00:00\",\"dayEndTime\":\"23:59\",\"week\":[]}]} type:0不限制，1限制 dayStartTime:开始时间 dayEndTime:结束时间 week:1,2,3,4,5,6,7
    
    AvailableTime   string `json:"available_time,omitempty"`
    

    // 使用说明
    
    Description   string `json:"description,omitempty"`
    

    // 绝对日期中的结束时间
    
    EndTime   string `json:"end_time,omitempty"`
    

    // 扩展字段
    
    ExtInfo   string `json:"ext_info,omitempty"`
    

    // 库存数量
    
    Inventory   int64 `json:"inventory,omitempty"`
    

    // 适用商品范围 值：ALL，PART_AVAILABLE，PART_UNAVAILABLE * 说明：全部商品可用，部分商品可用，部分商品不可用
    
    ItemCoverage   string `json:"item_coverage,omitempty"`
    

    // 最低消费
    
    MinCharge   int64 `json:"min_charge,omitempty"`
    

    // 券模板名称
    
    Name   string `json:"name,omitempty"`
    

    // 绝对日期中的开始时间
    
    StartTime   string `json:"start_time,omitempty"`
    

    // 相对日期中的当天是否有效
    
    TodayAvailable   bool `json:"today_available,omitempty"`
    

    // 券模板类型 * 值：GIFT、DISCOUNT、CASH      * 说明：礼品券、礼品券、礼品券
    
    Type   string `json:"type,omitempty"`
    

    // 使用场景DINE_IN 堂食,TAKE_OUT外卖
    
    UseCondition   string `json:"use_condition,omitempty"`
    

    // 每人限领
    
    UserLimit   int64 `json:"user_limit,omitempty"`
    

    // 券有效期类型 值： FIXED、RELATIVE      * 说明： 绝对时间，相对于用户领券时间
    
    ValidDateType   string `json:"valid_date_type,omitempty"`
    

    // 相对日期中的有效天数
    
    ValidDayCount   int64 `json:"valid_day_count,omitempty"`
    

    // 模版ID
    
    VoucherTemplateId   string `json:"voucher_template_id,omitempty"`
    

    // // 未使用     UNUSED,     // 使用中     USED,     // 无库存     NO_INVENTORY,     // 已失效     INVALID,     ;
    
    Status   string `json:"status,omitempty"`
    

    // 1
    
    ItemSelectedOpenInfoList  struct {
        ItemSelectedOpenInfo  []ItemSelectedOpenInfo `json:"item_selected_open_info,omitempty"`
    } `json:"item_selected_open_info_list,omitempty"`
    

    // 1
    
    ShopSelectedOpenInfoList  struct {
        ShopSelectedOpenInfo  []ShopSelectedOpenInfo `json:"shop_selected_open_info,omitempty"`
    } `json:"shop_selected_open_info_list,omitempty"`
    

    // 是否删除
    
    Deleted   bool `json:"deleted,omitempty"`
    

    // 面额
    
    Denomination   string `json:"denomination,omitempty"`
    

    // 更新人
    
    UpdateBy   string `json:"update_by,omitempty"`
    

    // 创建人
    
    CreateBy   string `json:"create_by,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

}
*/

// VoucherTemplateOpenInfo 
type VoucherTemplateOpenInfo struct {

    // {\"type\":0,\"settings\":[{\"dayStartTime\":\"00:00\",\"dayEndTime\":\"23:59\",\"week\":[]}]} type:0不限制，1限制 dayStartTime:开始时间 dayEndTime:结束时间 week:1,2,3,4,5,6,7
    AvailableTime   string `json:"available_time,omitempty"`

    // 使用说明
    Description   string `json:"description,omitempty"`

    // 绝对日期中的结束时间
    EndTime   string `json:"end_time,omitempty"`

    // 扩展字段
    ExtInfo   string `json:"ext_info,omitempty"`

    // 库存数量
    Inventory   int64 `json:"inventory,omitempty"`

    // 适用商品范围 值：ALL，PART_AVAILABLE，PART_UNAVAILABLE * 说明：全部商品可用，部分商品可用，部分商品不可用
    ItemCoverage   string `json:"item_coverage,omitempty"`

    // 最低消费
    MinCharge   int64 `json:"min_charge,omitempty"`

    // 券模板名称
    Name   string `json:"name,omitempty"`

    // 绝对日期中的开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 相对日期中的当天是否有效
    TodayAvailable   bool `json:"today_available,omitempty"`

    // 券模板类型 * 值：GIFT、DISCOUNT、CASH      * 说明：礼品券、礼品券、礼品券
    Type   string `json:"type,omitempty"`

    // 使用场景DINE_IN 堂食,TAKE_OUT外卖
    UseCondition   string `json:"use_condition,omitempty"`

    // 每人限领
    UserLimit   int64 `json:"user_limit,omitempty"`

    // 券有效期类型 值： FIXED、RELATIVE      * 说明： 绝对时间，相对于用户领券时间
    ValidDateType   string `json:"valid_date_type,omitempty"`

    // 相对日期中的有效天数
    ValidDayCount   int64 `json:"valid_day_count,omitempty"`

    // 模版ID
    VoucherTemplateId   string `json:"voucher_template_id,omitempty"`

    // // 未使用     UNUSED,     // 使用中     USED,     // 无库存     NO_INVENTORY,     // 已失效     INVALID,     ;
    Status   string `json:"status,omitempty"`

    // 1
    ItemSelectedOpenInfoList   []ItemSelectedOpenInfo `json:"item_selected_open_info_list,omitempty"`

    // 1
    ShopSelectedOpenInfoList   []ShopSelectedOpenInfo `json:"shop_selected_open_info_list,omitempty"`

    // 是否删除
    Deleted   bool `json:"deleted,omitempty"`

    // 面额
    Denomination   string `json:"denomination,omitempty"`

    // 更新人
    UpdateBy   string `json:"update_by,omitempty"`

    // 创建人
    CreateBy   string `json:"create_by,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
