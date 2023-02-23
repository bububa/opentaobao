package alihouse

// VideoDraftDto 结构体
type VideoDraftDto struct {
	// 二级栏目名称
	SubColumn string `json:"sub_column,omitempty" xml:"sub_column,omitempty"`
	// 一级栏目名称
	TopColumn string `json:"top_column,omitempty" xml:"top_column,omitempty"`
	// 视频来源
	VideoSource string `json:"video_source,omitempty" xml:"video_source,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 外部视频id
	OuterVideoId string `json:"outer_video_id,omitempty" xml:"outer_video_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 发布时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 封面图
	CoverImage string `json:"cover_image,omitempty" xml:"cover_image,omitempty"`
	// 图片信息
	ImgInfo string `json:"img_info,omitempty" xml:"img_info,omitempty"`
	// 视频下载地址
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 主播名称
	CompereName string `json:"compere_name,omitempty" xml:"compere_name,omitempty"`
	// 主播头像
	CompereHeadUrl string `json:"compere_head_url,omitempty" xml:"compere_head_url,omitempty"`
	// 作者
	Author string `json:"author,omitempty" xml:"author,omitempty"`
	// 视频介绍
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 比如楼盘下的货品id，小区小的房源id
	OuterSubId string `json:"outer_sub_id,omitempty" xml:"outer_sub_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 视频类型
	VideoType int64 `json:"video_type,omitempty" xml:"video_type,omitempty"`
	// 来源
	Original int64 `json:"original,omitempty" xml:"original,omitempty"`
	// 菜鸟城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 视频格式
	VideoFormat int64 `json:"video_format,omitempty" xml:"video_format,omitempty"`
	// 主播id
	CompereId int64 `json:"compere_id,omitempty" xml:"compere_id,omitempty"`
	// 视频长度
	VideoLength int64 `json:"video_length,omitempty" xml:"video_length,omitempty"`
	// 分辨率
	Resolution int64 `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// 1测试数据 0正常数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
