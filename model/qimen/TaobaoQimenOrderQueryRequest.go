package qimen

// TaobaoQimenOrderQueryRequest 
type TaobaoQimenOrderQueryRequest struct {
    // 姓名, string (50) , 必填
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 开始时间, string (19) , YYYY-MM-DD HH:MM:SS
    StartTime   string `json:"startTime,omitempty" xml:"startTime,omitempty"`
    // 结束时间, string (19) , YYYY-MM-DD HH:MM:SS
    EndTime   string `json:"endTime,omitempty" xml:"endTime,omitempty"`
    // 固定电话, string (50)
    Tel   string `json:"tel,omitempty" xml:"tel,omitempty"`
    // 移动电话, string (50) , 必填
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}
