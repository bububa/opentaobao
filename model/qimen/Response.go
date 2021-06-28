package qimen

// Response 
/* model for simplify = false
type Response struct {

    // 响应结果:success|failure
    
    Flag   string `json:"flag,omitempty"`
    

    // 响应码
    
    Code   string `json:"code,omitempty"`
    

    // 响应信息
    
    Message   string `json:"message,omitempty"`
    

    // 奇门仓储字段
    
    CombItems  struct {
        CombItem  []CombItem `json:"comb_item,omitempty"`
    } `json:"combItems,omitempty"`
    

    // 仓储系统入库单编码
    
    EntryOrderId   string `json:"entryOrderId,omitempty"`
    

    // 奇门仓储字段
    
    ExpressInfos  struct {
        ExpressInfo  []ExpressInfo `json:"express_info,omitempty"`
    } `json:"expressInfos,omitempty"`
    

    // 总行数，必填
    
    TotalCount   int64 `json:"totalCount,omitempty"`
    

    // 明细列表
    
    Items  *struct {
        Items  *Items `json:"items,omitempty"`
    } `json:"items,omitempty"`
    

    // 奇门仓储字段
    
    OrderCode   string `json:"orderCode,omitempty"`
    

    // 奇门仓储字段
    
    IsRetry   string `json:"isRetry,omitempty"`
    

    // 奇门仓储字段
    
    ItemInventories  struct {
        ItemInventory  []ItemInventory `json:"item_inventory,omitempty"`
    } `json:"itemInventories,omitempty"`
    

    // 商品映射关系
    
    ItemMappings  struct {
        ItemMapping  []ItemMapping `json:"item_mapping,omitempty"`
    } `json:"itemMappings,omitempty"`
    

    // 订单列表
    
    OrderLines  *struct {
        OrderLines  *OrderLines `json:"order_lines,omitempty"`
    } `json:"orderLines,omitempty"`
    

    // 订单收件人 ID, string (50)
    
    Oaid   string `json:"oaid,omitempty"`
    

    // 出库单号, string (50) , 必填
    
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`
    

    // 收货人信息
    
    ReceiverInfo  *struct {
        ReceiverInfo  *ReceiverInfo `json:"receiver_info,omitempty"`
    } `json:"receiverInfo,omitempty"`
    

    // 仓储系统退货单编码
    
    ReturnOrderId   string `json:"returnOrderId,omitempty"`
    

    // 仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传)
    
    ItemId   string `json:"itemId,omitempty"`
    

    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    
    CreateTime   string `json:"createTime,omitempty"`
    

    // 出库单仓储系统编码
    
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`
    

    // 奇门仓储字段
    
    OwnerCode   string `json:"ownerCode,omitempty"`
    

    // 奇门仓储字段
    
    OwnerName   string `json:"ownerName,omitempty"`
    

    // 奇门仓储字段
    
    WarehouseInfos  struct {
        WarehouseInfo  []WarehouseInfo `json:"warehouse_info,omitempty"`
    } `json:"warehouseInfos,omitempty"`
    

}
*/

// Response 
type Response struct {

    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty"`

    // 响应码
    Code   string `json:"code,omitempty"`

    // 响应信息
    Message   string `json:"message,omitempty"`

    // 奇门仓储字段
    CombItems   []CombItem `json:"combItems,omitempty"`

    // 仓储系统入库单编码
    EntryOrderId   string `json:"entryOrderId,omitempty"`

    // 奇门仓储字段
    ExpressInfos   []ExpressInfo `json:"expressInfos,omitempty"`

    // 总行数，必填
    TotalCount   int64 `json:"totalCount,omitempty"`

    // 明细列表
    Items   *Items `json:"items,omitempty"`

    // 奇门仓储字段
    OrderCode   string `json:"orderCode,omitempty"`

    // 奇门仓储字段
    IsRetry   string `json:"isRetry,omitempty"`

    // 奇门仓储字段
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty"`

    // 商品映射关系
    ItemMappings   []ItemMapping `json:"itemMappings,omitempty"`

    // 订单列表
    OrderLines   *OrderLines `json:"orderLines,omitempty"`

    // 订单收件人 ID, string (50)
    Oaid   string `json:"oaid,omitempty"`

    // 出库单号, string (50) , 必填
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty"`

    // 收货人信息
    ReceiverInfo   *ReceiverInfo `json:"receiverInfo,omitempty"`

    // 仓储系统退货单编码
    ReturnOrderId   string `json:"returnOrderId,omitempty"`

    // 仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传)
    ItemId   string `json:"itemId,omitempty"`

    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty"`

    // 出库单仓储系统编码
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty"`

    // 奇门仓储字段
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 奇门仓储字段
    OwnerName   string `json:"ownerName,omitempty"`

    // 奇门仓储字段
    WarehouseInfos   []WarehouseInfo `json:"warehouseInfos,omitempty"`

}
