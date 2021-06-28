package qimen

// Request 
/* model for simplify = false
type Request struct {

    // 奇门仓储字段
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 奇门仓储字段
    
    ItemId   string `json:"itemId,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

    // 奇门仓储字段
    
    ExpressCode   string `json:"expressCode,omitempty"`
    

    // 仓库编码，string(50)，必填
    
    WarehouseCode   string `json:"warehouseCode,omitempty"`
    

    // ERP 系统商品编码, string(50)，条件必填
    
    ItemCode   string `json:"itemCode,omitempty"`
    

    // 当前页，从 1 开始，必填
    
    Page   int64 `json:"page,omitempty"`
    

    // 每页条数(最多 100 条)，必填
    
    PageSize   int64 `json:"pageSize,omitempty"`
    

    // 奇门仓储字段
    
    OrderCode   string `json:"orderCode,omitempty"`
    

    // 奇门仓储字段
    
    OrderSource   string `json:"orderSource,omitempty"`
    

    // 奇门仓储字段
    
    ItemInventories  struct {
        ItemInventory  []ItemInventory `json:"item_inventory,omitempty"`
    } `json:"itemInventories,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    QueryType   string `json:"queryType,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    ShopItemId   string `json:"shopItemId,omitempty"`
    

    // 奇门仓储字段,C123,string(50),,
    
    SkuId   string `json:"skuId,omitempty"`
    

    // 姓名, string (50) , 必填
    
    Name   string `json:"name,omitempty"`
    

    // 开始时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    StartTime   string `json:"startTime,omitempty"`
    

    // 结束时间, string (19) , YYYY-MM-DD HH:MM:SS
    
    EndTime   string `json:"endTime,omitempty"`
    

    // 固定电话, string (50)
    
    Tel   string `json:"tel,omitempty"`
    

    // 移动电话, string (50) , 必填
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"totalPage,omitempty"`
    

    // 当前页(从1开始)
    
    CurrentPage   int64 `json:"currentPage,omitempty"`
    

    // 单据列表
    
    Order  *struct {
        Order  *Order `json:"order,omitempty"`
    } `json:"order,omitempty"`
    

    // 商品列表
    
    Items  *struct {
        Items  *Items `json:"items,omitempty"`
    } `json:"items,omitempty"`
    

    // 奇门仓储字段
    
    MessageId   string `json:"messageId,omitempty"`
    

    // 奇门仓储字段
    
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`
    

    // 奇门仓储字段
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`
    

    // 奇门仓储字段
    
    OrderType   string `json:"orderType,omitempty"`
    

    // 奇门仓储字段
    
    LogisticsCode   string `json:"logisticsCode,omitempty"`
    

    // 奇门仓储字段
    
    MessageType   string `json:"messageType,omitempty"`
    

    // 奇门仓储字段
    
    MessageDesc   string `json:"messageDesc,omitempty"`
    

    // 奇门仓储字段
    
    CreateTime   string `json:"createTime,omitempty"`
    

    // 奇门仓储字段
    
    OrderLines  struct {
        OrderLine  []OrderLine `json:"order_line,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 奇门仓储字段
    
    Remark   string `json:"remark,omitempty"`
    

    // 订单收件人 ID, string (50)
    
    Oaid   string `json:"oaid,omitempty"`
    

    // 包裹列表
    
    Packages  *struct {
        Packages  *Packages `json:"packages,omitempty"`
    } `json:"packages,omitempty"`
    

    // add,update, 必填
    
    ActionType   string `json:"actionType,omitempty"`
    

    // 店铺
    
    Shop  *struct {
        Shop  *Shop `json:"shop,omitempty"`
    } `json:"shop,omitempty"`
    

    // 供应商编码 string (50), 必填
    
    SupplierCode   string `json:"supplierCode,omitempty"`
    

    // 供应商名称 string (200) , 必填
    
    SupplierName   string `json:"supplierName,omitempty"`
    

    // 电子邮箱, string (50)
    
    Email   string `json:"email,omitempty"`
    

    // 国家二字码，string（50）
    
    CountryCode   string `json:"countryCode,omitempty"`
    

    // 省份, string (50)
    
    Province   string `json:"province,omitempty"`
    

    // 城市, string (50)
    
    City   string `json:"city,omitempty"`
    

    // 区域, string (50)
    
    Area   string `json:"area,omitempty"`
    

    // 村镇, string (50)
    
    Town   string `json:"town,omitempty"`
    

    // 详细地址, string (200)
    
    DetailAddress   string `json:"detailAddress,omitempty"`
    

    // 是否有效, Y/N (默认为Y)
    
    IsValid   string `json:"isValid,omitempty"`
    

    // 货主名称，string（50）
    
    OwnerName   string `json:"ownerName,omitempty"`
    

    // 仓库信息
    
    WarehouseInfos  struct {
        WarehouseInfos  []WarehouseInfos `json:"warehouse_infos,omitempty"`
    } `json:"warehouseInfos,omitempty"`
    

}
*/

// Request 
type Request struct {

    // 奇门仓储字段
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 奇门仓储字段
    ItemId   string `json:"itemId,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

    // 奇门仓储字段
    ExpressCode   string `json:"expressCode,omitempty"`

    // 仓库编码，string(50)，必填
    WarehouseCode   string `json:"warehouseCode,omitempty"`

    // ERP 系统商品编码, string(50)，条件必填
    ItemCode   string `json:"itemCode,omitempty"`

    // 当前页，从 1 开始，必填
    Page   int64 `json:"page,omitempty"`

    // 每页条数(最多 100 条)，必填
    PageSize   int64 `json:"pageSize,omitempty"`

    // 奇门仓储字段
    OrderCode   string `json:"orderCode,omitempty"`

    // 奇门仓储字段
    OrderSource   string `json:"orderSource,omitempty"`

    // 奇门仓储字段
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    QueryType   string `json:"queryType,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    ShopItemId   string `json:"shopItemId,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    SkuId   string `json:"skuId,omitempty"`

    // 姓名, string (50) , 必填
    Name   string `json:"name,omitempty"`

    // 开始时间, string (19) , YYYY-MM-DD HH:MM:SS
    StartTime   string `json:"startTime,omitempty"`

    // 结束时间, string (19) , YYYY-MM-DD HH:MM:SS
    EndTime   string `json:"endTime,omitempty"`

    // 固定电话, string (50)
    Tel   string `json:"tel,omitempty"`

    // 移动电话, string (50) , 必填
    Mobile   string `json:"mobile,omitempty"`

    // 总页数
    TotalPage   int64 `json:"totalPage,omitempty"`

    // 当前页(从1开始)
    CurrentPage   int64 `json:"currentPage,omitempty"`

    // 单据列表
    Order   *Order `json:"order,omitempty"`

    // 商品列表
    Items   *Items `json:"items,omitempty"`

    // 奇门仓储字段
    MessageId   string `json:"messageId,omitempty"`

    // 奇门仓储字段
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`

    // 奇门仓储字段
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`

    // 奇门仓储字段
    OrderType   string `json:"orderType,omitempty"`

    // 奇门仓储字段
    LogisticsCode   string `json:"logisticsCode,omitempty"`

    // 奇门仓储字段
    MessageType   string `json:"messageType,omitempty"`

    // 奇门仓储字段
    MessageDesc   string `json:"messageDesc,omitempty"`

    // 奇门仓储字段
    CreateTime   string `json:"createTime,omitempty"`

    // 奇门仓储字段
    OrderLines   []OrderLine `json:"orderLines,omitempty"`

    // 奇门仓储字段
    Remark   string `json:"remark,omitempty"`

    // 订单收件人 ID, string (50)
    Oaid   string `json:"oaid,omitempty"`

    // 包裹列表
    Packages   *Packages `json:"packages,omitempty"`

    // add,update, 必填
    ActionType   string `json:"actionType,omitempty"`

    // 店铺
    Shop   *Shop `json:"shop,omitempty"`

    // 供应商编码 string (50), 必填
    SupplierCode   string `json:"supplierCode,omitempty"`

    // 供应商名称 string (200) , 必填
    SupplierName   string `json:"supplierName,omitempty"`

    // 电子邮箱, string (50)
    Email   string `json:"email,omitempty"`

    // 国家二字码，string（50）
    CountryCode   string `json:"countryCode,omitempty"`

    // 省份, string (50)
    Province   string `json:"province,omitempty"`

    // 城市, string (50)
    City   string `json:"city,omitempty"`

    // 区域, string (50)
    Area   string `json:"area,omitempty"`

    // 村镇, string (50)
    Town   string `json:"town,omitempty"`

    // 详细地址, string (200)
    DetailAddress   string `json:"detailAddress,omitempty"`

    // 是否有效, Y/N (默认为Y)
    IsValid   string `json:"isValid,omitempty"`

    // 货主名称，string（50）
    OwnerName   string `json:"ownerName,omitempty"`

    // 仓库信息
    WarehouseInfos   []WarehouseInfos `json:"warehouseInfos,omitempty"`

}
