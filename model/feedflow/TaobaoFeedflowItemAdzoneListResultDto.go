package feedflow

// TaobaoFeedflowItemAdzoneListResultDTO 
type TaobaoFeedflowItemAdzoneListResultDTO struct {
    // 成功
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 调用是否成功,true-成功，false-失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 广告位列表
    AdzoneList   []AdzoneDTO `json:"adzone_list,omitempty" xml:"adzone_list>adzone_dto,omitempty"`
}
