package product

// HscodeAuditInfo 
type HscodeAuditInfo struct {
    // 商品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // SKU的ID
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 商品或SKU使用的HS海关代码
    Hscode   string `json:"hscode,omitempty" xml:"hscode,omitempty"`
    // hscode信息当前审核状态，HISTORY_ITEM：历史已上架商品，REJECT：审核未通过，AUDITING：审核中，PASS：审核通过，ERROR：获取审核状态异常
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // hscode信息当前审核状态的具体说明
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
}
