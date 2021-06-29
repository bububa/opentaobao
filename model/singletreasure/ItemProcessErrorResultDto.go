package singletreasure

// ItemProcessErrorResultDTO 
type ItemProcessErrorResultDTO struct {
    // activityId
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 处理失败的item和sku: map结构为:itemId: skuId-错误信息,所有返回的外层key是itemId，里面的key是skuId，商品级别的skuId为-1
    SkuIdMap   string `json:"sku_id_map,omitempty" xml:"sku_id_map,omitempty"`
}
