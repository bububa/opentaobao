package qimen

// Request 
type Request struct {

    // 奇门仓储字段
    
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    

    // 奇门仓储字段
    
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    

    // 扩展属性
    
    ExtendProps   *Map `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
    

    // 奇门仓储字段
    
    ExpressCode   string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
    

    // 仓库编码，string(50)，必填
    
    WarehouseCode   string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
    

    // ERP 系统商品编码, string(50)，条件必填
    
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
    

    // 当前页，从 1 开始，必填
    
    Page   int64 `json:"page,omitempty" xml:"page,omitempty"`
    

    // 每页条数(最多 100 条)，必填
    
    PageSize   int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
    

    // 奇门仓储字段
    
    OrderCode   string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
    

    // 奇门仓储字段
    
    OrderSource   string `json:"orderSource,omitempty" xml:"orderSource,omitempty"`
    

    // 奇门仓储字段
    
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    QueryType   string `json:"queryType,omitempty" xml:"queryType,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    ShopItemId   string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    SkuId   string `json:"skuId,omitempty" xml:"skuId,omitempty"`
    

    // 姓名, string (50) , 必填
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 开始时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    StartTime   string `json:"startTime,omitempty" xml:"startTime,omitempty"`
    

    // 结束时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    EndTime   string `json:"endTime,omitempty" xml:"endTime,omitempty"`
    

    // 固定电话, string (50)
    
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    

    // 移动电话, string (50) , 必填
    
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
    

    // 当前页(从1开始)
    
    CurrentPage   int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
    

    // 单据列表
    
    Order   *Order `json:"order,omitempty" xml:"order,omitempty"`
    

    // 商品列表
    
    Items   *Items `json:"items,omitempty" xml:"items,omitempty"`
    

    // 奇门仓储字段
    
    MessageId   string `json:"messageId,omitempty" xml:"messageId,omitempty"`
    

    // 奇门仓储字段
    
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    

    // 奇门仓储字段
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    

    // 奇门仓储字段
    
    OrderType   string `json:"orderType,omitempty" xml:"orderType,omitempty"`
    

    // 奇门仓储字段
    
    LogisticsCode   string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
    

    // 奇门仓储字段
    
    MessageType   string `json:"messageType,omitempty" xml:"messageType,omitempty"`
    

    // 奇门仓储字段
    
    MessageDesc   string `json:"messageDesc,omitempty" xml:"messageDesc,omitempty"`
    

    // 奇门仓储字段
    
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    

    // 奇门仓储字段
    
    OrderLines   []OrderLine `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    

    // 奇门仓储字段
    
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    

    // 订单收件人 ID, string (50)
    
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    

    // 包裹列表
    
    Packages   *Packages `json:"packages,omitempty" xml:"packages,omitempty"`
    

    // add,update, 必填
    
    ActionType   string `json:"actionType,omitempty" xml:"actionType,omitempty"`
    

    // 店铺
    
    Shop   *Shop `json:"shop,omitempty" xml:"shop,omitempty"`
    

    // 供应商编码 string (50), 必填
    
    SupplierCode   string `json:"supplierCode,omitempty" xml:"supplierCode,omitempty"`
    

    // 供应商名称 string (200) , 必填
    
    SupplierName   string `json:"supplierName,omitempty" xml:"supplierName,omitempty"`
    

    // 电子邮箱, string (50)
    
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    

    // 国家二字码，string（50）
    
    CountryCode   string `json:"countryCode,omitempty" xml:"countryCode,omitempty"`
    

    // 省份, string (50)
    
    Province   string `json:"province,omitempty" xml:"province,omitempty"`
    

    // 城市, string (50)
    
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    

    // 区域, string (50)
    
    Area   string `json:"area,omitempty" xml:"area,omitempty"`
    

    // 村镇, string (50)
    
    Town   string `json:"town,omitempty" xml:"town,omitempty"`
    

    // 详细地址, string (200)
    
    DetailAddress   string `json:"detailAddress,omitempty" xml:"detailAddress,omitempty"`
    

    // 是否有效, Y/N (默认为Y)
    
    IsValid   string `json:"isValid,omitempty" xml:"isValid,omitempty"`
    

    // 货主名称，string（50）
    
    OwnerName   string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
    

    // 仓库信息
    
    WarehouseInfos   []WarehouseInfos `json:"warehouseInfos,omitempty" xml:"warehouseInfos,omitempty"`
    

}
