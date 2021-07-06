package ascpqcc

// CancelSampleRelationData 结构体
type CancelSampleRelationData struct {
	// 商品编号列表
	BizItemIds []string `json:"biz_item_ids,omitempty" xml:"biz_item_ids>string,omitempty"`
	// 商品编号
	BizItemId string `json:"biz_item_id,omitempty" xml:"biz_item_id,omitempty"`
}
