package xiami

// AlibabaXiamiApiMobileFigureimageGetStruct 结构体
type AlibabaXiamiApiMobileFigureimageGetStruct struct {
	// pic_url_xmusic
	PicUrlXmusic string `json:"pic_url_xmusic,omitempty" xml:"pic_url_xmusic,omitempty"`
	// 内容
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// pic_url_yasha
	PicUrlYasha string `json:"pic_url_yasha,omitempty" xml:"pic_url_yasha,omitempty"`
	// 简介
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// pic_url_newindex
	PicUrlNewindex string `json:"pic_url_newindex,omitempty" xml:"pic_url_newindex,omitempty"`
	// 图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 链接地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 结束时间
	GmtEnd int64 `json:"gmt_end,omitempty" xml:"gmt_end,omitempty"`
	// 开始时间
	GmtPub int64 `json:"gmt_pub,omitempty" xml:"gmt_pub,omitempty"`
}
