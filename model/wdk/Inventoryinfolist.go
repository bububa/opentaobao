package wdk

// Inventoryinfolist 
type Inventoryinfolist struct {

    // 库存单位
    StorageUnit   string `json:"storage_unit,omitempty"`

    // 实物总量
    RealInvent   string `json:"real_invent,omitempty"`

    // 商品编码
    ItemCode   string `json:"item_code,omitempty"`

    // 货位库存详情
    CabinetInventInfos   []Cabinetinventinfos `json:"cabinet_invent_infos,omitempty"`

}
