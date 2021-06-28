package wms

// CainiaoInventoryProfitlossProfitlossinfo 
/* model for simplify = false
type CainiaoInventoryProfitlossProfitlossinfo struct {

    // 仓库编码
    
    StoreCode   string `json:"store_code,omitempty"`
    

    // 仓库订单编码
    
    CnOrderCode   string `json:"cn_order_code,omitempty"`
    

    // 订单类型： 701 盘点出库 702 盘点入库
    
    OrderType   int64 `json:"order_type,omitempty"`
    

    // 商品信息列表
    
    OrderItemList  struct {
        CainiaoInventoryProfitlossOrderitemlist  []CainiaoInventoryProfitlossOrderitemlist `json:"cainiao_inventory_profitloss_orderitemlist,omitempty"`
    } `json:"order_item_list,omitempty"`
    

    // 备注
    
    Remark   string `json:"remark,omitempty"`
    

    // 单据生成时间
    
    CreatedTime   string `json:"created_time,omitempty"`
    

}
*/

// CainiaoInventoryProfitlossProfitlossinfo 
type CainiaoInventoryProfitlossProfitlossinfo struct {

    // 仓库编码
    StoreCode   string `json:"store_code,omitempty"`

    // 仓库订单编码
    CnOrderCode   string `json:"cn_order_code,omitempty"`

    // 订单类型： 701 盘点出库 702 盘点入库
    OrderType   int64 `json:"order_type,omitempty"`

    // 商品信息列表
    OrderItemList   []CainiaoInventoryProfitlossOrderitemlist `json:"order_item_list,omitempty"`

    // 备注
    Remark   string `json:"remark,omitempty"`

    // 单据生成时间
    CreatedTime   string `json:"created_time,omitempty"`

}
