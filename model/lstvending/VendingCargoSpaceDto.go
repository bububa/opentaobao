package lstvending

// VendingCargoSpaceDto 
type VendingCargoSpaceDto struct {
    // 修改时间
    GmtModified   int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 货道商品数量
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 创建时间
    GmtCreate   int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 商品折扣价
    DiscountPrice   int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
    // 外部商品ID
    ExternalGoodsId   string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
    // 货道编号，从1开始
    CargoRoadNo   int64 `json:"cargo_road_no,omitempty" xml:"cargo_road_no,omitempty"`
    // 厂商设备唯一编码
    EquipmentCode   string `json:"equipment_code,omitempty" xml:"equipment_code,omitempty"`
    // 货道类型：1普通货道，2VIP货道
    BizType   int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
    // 货架编号，从1开始
    ShelfNo   int64 `json:"shelf_no,omitempty" xml:"shelf_no,omitempty"`
    // 商品单价，单位：分
    Price   int64 `json:"price,omitempty" xml:"price,omitempty"`
    // 货道状态：1正常，2故障
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 厂商货道ID
    ExternalId   string `json:"external_id,omitempty" xml:"external_id,omitempty"`
    // 商品ID
    GoodsId   int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
    // 供应商编码
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // 货道ID
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
}
