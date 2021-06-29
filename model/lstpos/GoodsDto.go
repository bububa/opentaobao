package lstpos

// GoodsDTO 
type GoodsDTO struct {
    // 规格
    Spec   string `json:"spec,omitempty" xml:"spec,omitempty"`
    // 简称
    ShotTitle   string `json:"shot_title,omitempty" xml:"shot_title,omitempty"`
    // 零售价，单位为分
    ReservePrice   int64 `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
    // 商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 商品条码,唯一
    BarCode   string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
    // goodsId
    GoodsId   int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
    // originalPrice
    OriginalPrice   int64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
    // categoryId
    CategoryId   string `json:"category_id,omitempty" xml:"category_id,omitempty"`
    // ISV名称
    Source   string `json:"source,omitempty" xml:"source,omitempty"`
    // isvGoodsId
    IsvGoodsId   string `json:"isv_goods_id,omitempty" xml:"isv_goods_id,omitempty"`
    // shortcutCode
    ShortcutCode   string `json:"shortcut_code,omitempty" xml:"shortcut_code,omitempty"`
    // 称重单位，比如：按件（c），按重量（w：斤）
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 商品状态：ON上架、OFF下架、DELETE已删除
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 计价方式：按数量（c）、按重量（w）
    PricingMode   string `json:"pricing_mode,omitempty" xml:"pricing_mode,omitempty"`
    // 商品所属小店用户ID
    UserId   int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
    // 创建时间，精确到毫秒
    GmtCreate   int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间，精确到毫秒
    GmtModified   int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 自营叶子类目自定义编号(ISV子定义类目Id)
    IsvCategoryId   string `json:"isv_category_id,omitempty" xml:"isv_category_id,omitempty"`
}
