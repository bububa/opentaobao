package xiami

// AlibabaXiamiApiMvMusiclistGetStruct 结构体
type AlibabaXiamiApiMvMusiclistGetStruct struct {
	// mv列表
	Mvs []AlibabaXiamiApiMvMusiclistGetStruct `json:"mvs,omitempty" xml:"mvs>alibaba_xiami_api_mv_musiclist_get_struct,omitempty"`
	// mv关联艺人
	Artists []AlibabaXiamiApiMvMusiclistGetStruct `json:"artists,omitempty" xml:"artists>alibaba_xiami_api_mv_musiclist_get_struct,omitempty"`
	// MV封面
	MvCover string `json:"mv_cover,omitempty" xml:"mv_cover,omitempty"`
	// MV标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// MV ID
	MvId string `json:"mv_id,omitempty" xml:"mv_id,omitempty"`
	// 艺人名称
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 视频地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// MV 播放信息
	Video *AlibabaXiamiApiMvMusiclistGetStruct `json:"video,omitempty" xml:"video,omitempty"`
	// 艺人ID, BIGINT类型
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// 过期时间，为0时表示不过期，如果时间小于当前时间戳，则需要重新请求接口
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// 是否有下一页
	More bool `json:"more,omitempty" xml:"more,omitempty"`
}
