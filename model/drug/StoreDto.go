package drug

// StoreDto 
type StoreDto struct {

    // shopId
    ShopId   string `json:"shop_id,omitempty"`

    // storeId
    StoreId   string `json:"store_id,omitempty"`

    // storeName
    StoreName   string `json:"store_name,omitempty"`

    // picURL
    PicUrl   string `json:"pic_url,omitempty"`

    // businessStatus
    BusinessStatus   string `json:"business_status,omitempty"`

    // orderCount
    OrderCount   int64 `json:"order_count,omitempty"`

    // orderCountDesc
    OrderCountDesc   string `json:"order_count_desc,omitempty"`

    // minimumAccount
    MinimumAccount   int64 `json:"minimum_account,omitempty"`

    // deliveryAccount
    DeliveryAccount   int64 `json:"delivery_account,omitempty"`

    // notice
    Notice   string `json:"notice,omitempty"`

    // deliveryTime
    DeliveryTime   string `json:"delivery_time,omitempty"`

    // address
    Address   string `json:"address,omitempty"`

    // tele
    Tele   string `json:"tele,omitempty"`

    // latitude
    Latitude   string `json:"latitude,omitempty"`

    // longitude
    Longitude   string `json:"longitude,omitempty"`

    // earliestTimeOfDelivery
    EarliestTimeOfDelivery   string `json:"earliest_time_of_delivery,omitempty"`

    // punctualityRate
    PunctualityRate   string `json:"punctuality_rate,omitempty"`

    // dpavgscore
    Dpavgscore   string `json:"dpavgscore,omitempty"`

    // areas
    Areas   string `json:"areas,omitempty"`

    // sellerNick
    SellerNick   string `json:"seller_nick,omitempty"`

    // deliveryType
    DeliveryType   int64 `json:"delivery_type,omitempty"`

    // deliveryTypeDesc
    DeliveryTypeDesc   string `json:"delivery_type_desc,omitempty"`

}
