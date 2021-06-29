package tmallnr

// NrInventoryUpdateReqDTO 
type NrInventoryUpdateReqDTO struct {
    // 1表示全量，2表示增量
    CheckMode   int64 `json:"check_mode,omitempty" xml:"check_mode,omitempty"`
    // 商家的sellerID，如果是零售商需要给品牌的sellerId
    OwnerId   int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
    // 更新库存的列表值
    DetailList   []NrInventoryCheckDetailDTO `json:"detail_list,omitempty" xml:"detail_list>nr_inventory_check_detail_dto,omitempty"`
    // 定时送dss，jsd极速达
    BizIdentity   string `json:"biz_identity,omitempty" xml:"biz_identity,omitempty"`
    // 幂等值
    OrderId   string `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 门店编号
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 默认为6：门店库存，2：电商库存
    StoreType   int64 `json:"store_type,omitempty" xml:"store_type,omitempty"`
}
