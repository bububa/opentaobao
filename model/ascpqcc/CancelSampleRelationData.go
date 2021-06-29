package ascpqcc

// CancelSampleRelationData 
type CancelSampleRelationData struct {
    // 商品编号
    BizItemId   string `json:"biz_item_id,omitempty" xml:"biz_item_id,omitempty"`
    // 商品编号列表
    BizItemIds   []string `json:"biz_item_ids,omitempty" xml:"biz_item_ids>string,omitempty"`
}
