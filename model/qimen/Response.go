package qimen

// Response 
type Response struct {
    // 响应结果:success|failure
    Flag   string `json:"flag,omitempty" xml:"flag,omitempty"`
    // 响应码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // 响应信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 奇门仓储字段
    CombItems   []CombItem `json:"combItems,omitempty" xml:"combItems>comb_item,omitempty"`
    // 仓储系统入库单编码
    EntryOrderId   string `json:"entryOrderId,omitempty" xml:"entryOrderId,omitempty"`
    // 奇门仓储字段
    ExpressInfos   []ExpressInfo `json:"expressInfos,omitempty" xml:"expressInfos>express_info,omitempty"`
    // 总行数，必填
    TotalCount   int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
    // 明细列表
    Items   *Items `json:"items,omitempty" xml:"items,omitempty"`
    // 奇门仓储字段
    OrderCode   string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
    // 奇门仓储字段
    IsRetry   string `json:"isRetry,omitempty" xml:"isRetry,omitempty"`
    // 奇门仓储字段
    ItemInventories   []ItemInventory `json:"itemInventories,omitempty" xml:"itemInventories>item_inventory,omitempty"`
    // 商品映射关系
    ItemMappings   []ItemMapping `json:"itemMappings,omitempty" xml:"itemMappings>item_mapping,omitempty"`
    // 订单列表
    OrderLines   *OrderLines `json:"orderLines,omitempty" xml:"orderLines,omitempty"`
    // 订单收件人 ID, string (50)
    Oaid   string `json:"oaid,omitempty" xml:"oaid,omitempty"`
    // 出库单号, string (50) , 必填
    DeliveryOrderCode   string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
    // 收货人信息
    ReceiverInfo   *ReceiverInfo `json:"receiverInfo,omitempty" xml:"receiverInfo,omitempty"`
    // 仓储系统退货单编码
    ReturnOrderId   string `json:"returnOrderId,omitempty" xml:"returnOrderId,omitempty"`
    // 仓储系统商品Id(当这个字段不为空的时候;所有erp传输的时候都碰到itemid必传)
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    // 订单创建时间(YYYY-MM-DD HH:MM:SS)
    CreateTime   string `json:"createTime,omitempty" xml:"createTime,omitempty"`
    // 出库单仓储系统编码
    DeliveryOrderId   string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
    // 奇门仓储字段
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段
    OwnerName   string `json:"ownerName,omitempty" xml:"ownerName,omitempty"`
    // 奇门仓储字段
    WarehouseInfos   []WarehouseInfo `json:"warehouseInfos,omitempty" xml:"warehouseInfos>warehouse_info,omitempty"`
}
