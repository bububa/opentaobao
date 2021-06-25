package train

// StationPassInfo 
type StationPassInfo struct {

    // 到达时间
    Atime   string `json:"atime,omitempty"`

    // 驶离时间
    Ltime   string `json:"ltime,omitempty"`

    // 当前站点索引，从1开始
    Sid   int64 `json:"sid,omitempty"`

    // 站点名称
    Snm   string `json:"snm,omitempty"`

    // 停留时间分钟
    Stime   int64 `json:"stime,omitempty"`

}
