package alicom

// ActivityInfoResponseDto 
type ActivityInfoResponseDto struct {

    // 活动ID
    ActivityId   string `json:"activity_id,omitempty"`

    // 活动名称
    ActivityName   string `json:"activity_name,omitempty"`

    // 礼包数据集
    ActivityGiftInfos   *ActivityGiftInfoDTOs `json:"activity_gift_infos,omitempty"`

}
