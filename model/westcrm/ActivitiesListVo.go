package westcrm

// ActivitiesListVo 
type ActivitiesListVo struct {
    // 活动有效期-结束时间
    TimeEnd   string `json:"time_end,omitempty" xml:"time_end,omitempty"`
    // 头图地址
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 活动状态：1-未开始；2-进行中；3-已结束
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 目标客群结构体
    Tags   *ThirdTagsVo `json:"tags,omitempty" xml:"tags,omitempty"`
    // 活动是否在活动列表中展示
    IsDisplay   int64 `json:"is_display,omitempty" xml:"is_display,omitempty"`
    // 活动名称
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 活动类型：0商场活动，1互动活动
    ActivityType   int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
    // 排序权重
    Sort   int64 `json:"sort,omitempty" xml:"sort,omitempty"`
    // 园区id
    CampusId   int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
    // 活动有效期-开始时间
    TimeBegin   string `json:"time_begin,omitempty" xml:"time_begin,omitempty"`
    // 活动id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
