package interact

// ActivityWriteResult 
type ActivityWriteResult struct {
    // 报名成功的id
    ActivityId   int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
    // 活动页面h5链接
    H5Url   string `json:"h5_url,omitempty" xml:"h5_url,omitempty"`
}
