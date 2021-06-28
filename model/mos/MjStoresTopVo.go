package mos

// MjStoresTopVo 
/* model for simplify = false
type MjStoresTopVo struct {

    // 默认选中的专柜，0为全部品牌
    
    StoreDefault   int64 `json:"store_default,omitempty"`
    

    // storeInfoList
    
    StoreInfoList  struct {
        StoreInfo  []StoreInfo `json:"store_info,omitempty"`
    } `json:"store_info_list,omitempty"`
    

    // 版本
    
    Version   int64 `json:"version,omitempty"`
    

    // 1：普通 2：刷脸
    
    ScreenType   int64 `json:"screen_type,omitempty"`
    

    // 商场id
    
    MallId   int64 `json:"mall_id,omitempty"`
    

    // 外部门店号
    
    OutMallId   string `json:"out_mall_id,omitempty"`
    

}
*/

// MjStoresTopVo 
type MjStoresTopVo struct {

    // 默认选中的专柜，0为全部品牌
    StoreDefault   int64 `json:"store_default,omitempty"`

    // storeInfoList
    StoreInfoList   []StoreInfo `json:"store_info_list,omitempty"`

    // 版本
    Version   int64 `json:"version,omitempty"`

    // 1：普通 2：刷脸
    ScreenType   int64 `json:"screen_type,omitempty"`

    // 商场id
    MallId   int64 `json:"mall_id,omitempty"`

    // 外部门店号
    OutMallId   string `json:"out_mall_id,omitempty"`

}
