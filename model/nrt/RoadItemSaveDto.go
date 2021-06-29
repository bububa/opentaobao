package nrt

// RoadItemSaveDto 
type RoadItemSaveDto struct {
    // 货道编号，若无跟商品编码保持一致
    RoadIds   string `json:"road_ids,omitempty" xml:"road_ids,omitempty"`
    // true 可售 false 不可售
    Available   bool `json:"available,omitempty" xml:"available,omitempty"`
    // 库存数，全量
    Inventory   int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
    // 容量
    Volume   int64 `json:"volume,omitempty" xml:"volume,omitempty"`
    // 商品编码
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 条码
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
}
