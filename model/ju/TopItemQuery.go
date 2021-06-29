package ju

// TopItemQuery 
type TopItemQuery struct {
    // 页码,必传
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 一页大小,必传
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 媒体pid,必传
    Pid   string `json:"pid,omitempty" xml:"pid,omitempty"`
    // 是否包邮,可不传
    Postage   bool `json:"postage,omitempty" xml:"postage,omitempty"`
    // 状态，预热：1，正在进行中：2,可不传
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 淘宝类目id,可不传
    TaobaoCategoryId   int64 `json:"taobao_category_id,omitempty" xml:"taobao_category_id,omitempty"`
    // 搜索关键词,可不传
    Word   string `json:"word,omitempty" xml:"word,omitempty"`
}
