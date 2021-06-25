package trade

// TradeFlowGoodsDetailModel 
type TradeFlowGoodsDetailModel struct {

    // 货道剩余商品数量
    RemainingQuantity   int64 `json:"remaining_quantity,omitempty"`

    // 货道业务类型：1普通；2推广实付金额
    BizType   int64 `json:"biz_type,omitempty"`

    // 实付金额(单位:分)
    ActualAmount   int64 `json:"actual_amount,omitempty"`

    // 商品数量
    Count   int64 `json:"count,omitempty"`

    // 货架编码，方向：从上到下，编码：从1开始
    ShelfNo   int64 `json:"shelf_no,omitempty"`

    // 交易总额(单位:分)
    TotalAmount   int64 `json:"total_amount,omitempty"`

    // 商品最小销售单位，如：包、盒、袋
    Unit   string `json:"unit,omitempty"`

    // 商品单价(单位:分)
    Price   int64 `json:"price,omitempty"`

    // 外部系统商品ID
    ExternalGoodsId   string `json:"external_goods_id,omitempty"`

    // 商品标题
    GoodsTitle   string `json:"goods_title,omitempty"`

    // 商品分类
    Category   string `json:"category,omitempty"`

    // 商品条码
    Barcode   string `json:"barcode,omitempty"`

    // 货道编码，方向：从左到右，编码：从1开始
    CargoRoadNo   int64 `json:"cargo_road_no,omitempty"`

}
