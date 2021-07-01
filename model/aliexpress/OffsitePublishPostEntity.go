package aliexpress

// OffsitePublishPostEntity 结构体
type OffsitePublishPostEntity struct {
	// 帖子的summary
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 图片列表，类型为1时不可为空
	ImageInfos []OffsitePostImageVO `json:"image_infos,omitempty" xml:"image_infos>offsite_post_image_vo,omitempty"`
	// 帖子类型：1-图文，2-视频
	PostType string `json:"post_type,omitempty" xml:"post_type,omitempty"`
	// 商品id
	ProductIds []int64 `json:"product_ids,omitempty" xml:"product_ids>int64,omitempty"`
	// 视频参数，类型为2时不可为空
	VideoInfo *OffsitePostVideoVO `json:"video_info,omitempty" xml:"video_info,omitempty"`
	// 帖子来源，import-第三方导入，接入前请咨询API负责人
	Origin string `json:"origin,omitempty" xml:"origin,omitempty"`
	// 语言地区
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 帖子拓展信息，垂类，形式等，介入前请咨询API负责人
	ExtendsInfo string `json:"extends_info,omitempty" xml:"extends_info,omitempty"`
	// 帖子话题
	Hashtag []string `json:"hashtag,omitempty" xml:"hashtag>string,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
}
