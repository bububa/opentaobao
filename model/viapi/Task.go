package viapi

// Task 
type Task struct {
    // 数据ID
    DataId   string `json:"data_id,omitempty" xml:"data_id,omitempty"`
    // 待检测图像的URL。支持HTTP和HTTPS协议。当前仅支持上海地域的OSS链接
    ImageUrl   string `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 图片创建或编辑时间
    ImageTime   int64 `json:"image_time,omitempty" xml:"image_time,omitempty"`
    // 截帧频率
    Interval   int64 `json:"interval,omitempty" xml:"interval,omitempty"`
    // 最大截帧数量
    MaxFrames   int64 `json:"max_frames,omitempty" xml:"max_frames,omitempty"`
    // 待检测的内容
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
}
