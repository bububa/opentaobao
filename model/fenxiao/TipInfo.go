package fenxiao

// TipInfo 
type TipInfo struct {

    // 返回信息
    Info   string `json:"info,omitempty"`

    // 商品id
    ScItemId   string `json:"sc_item_id,omitempty"`

    // errorCode
    Errorcode   string `json:"errorcode,omitempty"`

    // errorMessage
    Errormessage   string `json:"errormessage,omitempty"`

    // scItemCode
    ScItemCode   string `json:"sc_item_code,omitempty"`

    // storeCode
    StoreCode   string `json:"store_code,omitempty"`

    // invStoreType
    InvStoreType   int64 `json:"inv_store_type,omitempty"`

    // skuId
    SkuId   int64 `json:"sku_id,omitempty"`

    // 1前端商品 2供应链货品
    ItemType   int64 `json:"item_type,omitempty"`

}
