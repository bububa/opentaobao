package wms

// CainiaoStockOutBillStockoutinfo 
type CainiaoStockOutBillStockoutinfo struct {

    // ERP订单号
    OrderCode   string `json:"order_code,omitempty"`

    // 仓库订单编码，LBX号
    CnOrderCode   string `json:"cn_order_code,omitempty"`

    // 单据类型 903 普通出库单 305 B2B出库单 901 退供出库单
    OrderType   int64 `json:"order_type,omitempty"`

    // 仓库订单完成时间
    ConfirmTime   string `json:"confirm_time,omitempty"`

    // 入库单状态
    Status   string `json:"status,omitempty"`

    // 包裹信息列表，包含每个包裹使用的快递信息
    PackageInfoList   []CainiaoStockOutBillPackageinfolist `json:"package_info_list,omitempty"`

    // 订单商品列表
    OrderItemList   []CainiaoStockOutBillOrderitemlist `json:"order_item_list,omitempty"`

}
