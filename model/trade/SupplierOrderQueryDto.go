package trade

// SupplierOrderQueryDto 
type SupplierOrderQueryDto struct {

    // 外部门店ID
    OuterStoreId   string `json:"outer_store_id,omitempty"`

    // 查询交易创建时间在2020-06-03的订单，只接受yyyy-MM-dd格式的字符串
    TradeCreateDate   string `json:"trade_create_date,omitempty"`

    // 分页游标
    PageIndex   int64 `json:"page_index,omitempty"`

    // 供货商身份标识，比如大润发就传DRF
    Supplier   string `json:"supplier,omitempty"`

    // 分页参数，页大小
    PageSize   int64 `json:"page_size,omitempty"`

}
