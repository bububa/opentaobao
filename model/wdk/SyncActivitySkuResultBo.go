package wdk

// SyncActivitySkuResultBo 
type SyncActivitySkuResultBo struct {

    // 商品池ID
    PoolId   int64 `json:"pool_id,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 活动Id
    PromotionId   string `json:"promotion_id,omitempty"`

    // 版本Id
    VersionId   int64 `json:"version_id,omitempty"`

}
