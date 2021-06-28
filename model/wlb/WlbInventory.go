package wlb

// WlbInventory 
/* model for simplify = false
type WlbInventory struct {

    // 商品ID
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 货主ID
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 仓库编码，关联到仓库类型服务的编码非托管库存(卖家自己管理的库存，物流宝不可见又称自有库存)的所在仓库编码: STORE_SYS_PRIVATE
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 可销库存数量(库存总数-拍下预扣数-占用数)
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 冻结(锁定)数量，用来跟踪库存的中间状态，比如前台销售了1件商品，这时lock加1，当商品出库的时候lock再减回去
    
    LockQuantity   int64 `json:"lock_quantity,omitempty"`
    

    // VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存
    
    Type   string `json:"type,omitempty"`
    

    // 买家拍下但是没有付款的预扣数量，超过一定时间没有付款会自动释放。为防止超卖，这部分库存不可售，也暂不能出库。
    
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`
    

    // 库存占用数，两部分组成：<br/>1.交易占用<br/>买家付款后在仓库发货出库前会占用库存。<br/>2.非交易出库占用<br/>出库单创建仍未出库时会将库存占用
    
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`
    

}
*/

// WlbInventory 
type WlbInventory struct {

    // 商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 货主ID
    UserId   int64 `json:"user_id,omitempty"`

    // 仓库编码，关联到仓库类型服务的编码非托管库存(卖家自己管理的库存，物流宝不可见又称自有库存)的所在仓库编码: STORE_SYS_PRIVATE
    StoreCode   string `json:"store_code,omitempty"`

    // 可销库存数量(库存总数-拍下预扣数-占用数)
    Quantity   int64 `json:"quantity,omitempty"`

    // 冻结(锁定)数量，用来跟踪库存的中间状态，比如前台销售了1件商品，这时lock加1，当商品出库的时候lock再减回去
    LockQuantity   int64 `json:"lock_quantity,omitempty"`

    // VENDIBLE--可销售库存<br/>FREEZE--冻结库存<br/>ONWAY--在途库存<br/>DEFECT--残次品库存
    Type   string `json:"type,omitempty"`

    // 买家拍下但是没有付款的预扣数量，超过一定时间没有付款会自动释放。为防止超卖，这部分库存不可售，也暂不能出库。
    ReserveQuantity   int64 `json:"reserve_quantity,omitempty"`

    // 库存占用数，两部分组成：<br/>1.交易占用<br/>买家付款后在仓库发货出库前会占用库存。<br/>2.非交易出库占用<br/>出库单创建仍未出库时会将库存占用
    OccupyQuantity   int64 `json:"occupy_quantity,omitempty"`

}
