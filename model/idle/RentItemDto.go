package idle

// RentItemDTO 
type RentItemDTO struct {
    // 地址信息
    Address   *AddressDTO `json:"address,omitempty" xml:"address,omitempty"`
    // 类目id
    CatId   int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
    // 至少5个字符
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 标签
    FeaturedTags   []string `json:"featured_tags,omitempty" xml:"featured_tags>string,omitempty"`
    // sku
    ItemSkuList   []ItemSkuDTO `json:"item_sku_list,omitempty" xml:"item_sku_list>item_sku_dto,omitempty"`
    // 商品图片
    Media   *MediaDTO `json:"media,omitempty" xml:"media,omitempty"`
    // 价格信息
    PriceInfo   *PriceDTO `json:"price_info,omitempty" xml:"price_info,omitempty"`
    // 类目属性对
    PropPairs   []ItemPvPairDTO `json:"prop_pairs,omitempty" xml:"prop_pairs>item_pv_pair_dto,omitempty"`
    // 库存，无sku信息时必选
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 标配信息
    StandardEquipments   []EquipmentDTO `json:"standard_equipments,omitempty" xml:"standard_equipments>equipment_dto,omitempty"`
    // 标题，最少一个5个字符
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 新旧程度，10表示全新，9表示九成新
    UsedLevel   int64 `json:"used_level,omitempty" xml:"used_level,omitempty"`
    // 物流模板id
    TemplateId   int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
    // true表示包邮，否则表示不包邮
    FreePostage   bool `json:"free_postage,omitempty" xml:"free_postage,omitempty"`
    // 要更新的商品id
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 状态，可选值为0（表示上架），-2（表示下架），不填默认为0
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
}
