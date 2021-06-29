package singletreasure

// ItemDetailInfoBatchCreateDTO 
type ItemDetailInfoBatchCreateDTO struct {
    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 商品列表
    ItemDetailInfo   []ItemDetailInfoCreateDTO `json:"item_detail_info,omitempty" xml:"item_detail_info>item_detail_info_create_dto,omitempty"`
}
