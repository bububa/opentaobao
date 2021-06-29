package nrt

// NrtItemSyncResultDTO 
type NrtItemSyncResultDTO struct {
    // 摊位商品ID
    SItemId   int64 `json:"s_item_id,omitempty" xml:"s_item_id,omitempty"`
    // 主商品ID
    MItemId   int64 `json:"m_item_id,omitempty" xml:"m_item_id,omitempty"`
}
