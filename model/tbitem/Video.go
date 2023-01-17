package tbitem

// Video 结构体
type Video struct {
	// 视频关联记录创建时间（格式：yyyy-MM-dd HH:mm:ss）
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 视频关联记录修改时间（格式：yyyy-MM-dd HH:mm:ss）
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// video的url连接地址。淘秀里视频记录里面存储的url地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 视频关联记录的id，和商品相对应
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// video的id，对应于视频在淘秀的存储记录
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
}
