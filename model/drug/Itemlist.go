package drug

// Itemlist 
type Itemlist struct {
    // rx
    Rx   bool `json:"rx,omitempty" xml:"rx,omitempty"`
    // backCate
    BackCate   int64 `json:"back_cate,omitempty" xml:"back_cate,omitempty"`
    // itemId
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // isO2O
    IsO2o   string `json:"is_o2o,omitempty" xml:"is_o2o,omitempty"`
    // itemName
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // listPicUrl
    ListPicUrl   string `json:"list_pic_url,omitempty" xml:"list_pic_url,omitempty"`
    // oriPrice
    OriPrice   string `json:"ori_price,omitempty" xml:"ori_price,omitempty"`
    // price
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // symptom
    Symptom   string `json:"symptom,omitempty" xml:"symptom,omitempty"`
    // quantity
    Quantity   string `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // deliveryTime
    DeliveryTime   string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
    // atLimit
    AtLimit   int64 `json:"at_limit,omitempty" xml:"at_limit,omitempty"`
    // deliveryType
    DeliveryType   int64 `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
    // deliveryTypeDesc
    DeliveryTypeDesc   string `json:"delivery_type_desc,omitempty" xml:"delivery_type_desc,omitempty"`
}
