package tmallgenie

import (
	"sync"
)

// RawSingleVideo 结构体
type RawSingleVideo struct {
	// 系统标签ID，取值参见文档说明
	TagIds []int64 `json:"tag_ids,omitempty" xml:"tag_ids>int64,omitempty"`
	// 演员
	ActorName []string `json:"actor_name,omitempty" xml:"actor_name>string,omitempty"`
	// 导演
	DirectorName []string `json:"director_name,omitempty" xml:"director_name>string,omitempty"`
	// 别名
	Alias []string `json:"alias,omitempty" xml:"alias>string,omitempty"`
	// 支撑的分辨率
	SupportDefinition []string `json:"support_definition,omitempty" xml:"support_definition>string,omitempty"`
	// 出品人
	ProducerName []string `json:"producer_name,omitempty" xml:"producer_name>string,omitempty"`
	// 上传者名称
	UploaderName []string `json:"uploader_name,omitempty" xml:"uploader_name>string,omitempty"`
	// 所属视频集ID
	AlbumId string `json:"album_id,omitempty" xml:"album_id,omitempty"`
	// 视频描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 播放链接，Map，key为support_definition的值
	PlayUrl string `json:"play_url,omitempty" xml:"play_url,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 视频ID,不同的专辑间也要保证唯一
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 视频封面(竖图296 * 440，根据搜索规则，提供竖图用户搜索时可优先搜索到此内容)
	VCoverUrl string `json:"v_cover_url,omitempty" xml:"v_cover_url,omitempty"`
	// 视频来源类型，PGC/UGC/OGC
	OupgcType string `json:"oupgc_type,omitempty" xml:"oupgc_type,omitempty"`
	// 视频封面(横图，图片尺寸是295 * 167)
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 操作类型ADD/DELETE/UPDATE
	Operation string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 扩展字段
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 发布时间，unix时间戳，单位毫秒
	ReleaseTime int64 `json:"release_time,omitempty" xml:"release_time,omitempty"`
	// 付费类型 0-免费、1-VIP免费、2-单点、3-用券
	ChargeType int64 `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// 点赞数
	LikeCount int64 `json:"like_count,omitempty" xml:"like_count,omitempty"`
	// 视频时长，单位:秒
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 是否独播 0-非独播 1-独播
	IsExclusive int64 `json:"is_exclusive,omitempty" xml:"is_exclusive,omitempty"`
	// 失效时间,unix时间戳,单位:毫秒,为空则永不失效
	ValidEndTime int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	// 转发量
	ForwardCount int64 `json:"forward_count,omitempty" xml:"forward_count,omitempty"`
	// 系统类目ID，取值范围见文档
	CommonCateId int64 `json:"common_cate_id,omitempty" xml:"common_cate_id,omitempty"`
	// 生效时间,unix时间戳,单位:毫秒,为空则立即生效
	ValidStartTime int64 `json:"valid_start_time,omitempty" xml:"valid_start_time,omitempty"`
	// 收藏量
	CollectCount int64 `json:"collect_count,omitempty" xml:"collect_count,omitempty"`
	// 评论量
	CommentCount int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty"`
	// 播放量
	PlayCount int64 `json:"play_count,omitempty" xml:"play_count,omitempty"`
	// 视频在所属视频集中的排序号
	OrderIndex int64 `json:"order_index,omitempty" xml:"order_index,omitempty"`
	// 内容质量评分
	ContentScore int64 `json:"content_score,omitempty" xml:"content_score,omitempty"`
	// 建议最高零售价(单位:分)
	SuggestMaxPrice int64 `json:"suggest_max_price,omitempty" xml:"suggest_max_price,omitempty"`
	// 建议最低零售价(单位:分)
	SuggestMinPrice int64 `json:"suggest_min_price,omitempty" xml:"suggest_min_price,omitempty"`
	// 成本价(单位:分)
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 试看秒数，若支持试看填1，否则填0
	AuditionSecond int64 `json:"audition_second,omitempty" xml:"audition_second,omitempty"`
}

var poolRawSingleVideo = sync.Pool{
	New: func() any {
		return new(RawSingleVideo)
	},
}

// GetRawSingleVideo() 从对象池中获取RawSingleVideo
func GetRawSingleVideo() *RawSingleVideo {
	return poolRawSingleVideo.Get().(*RawSingleVideo)
}

// ReleaseRawSingleVideo 释放RawSingleVideo
func ReleaseRawSingleVideo(v *RawSingleVideo) {
	v.TagIds = v.TagIds[:0]
	v.ActorName = v.ActorName[:0]
	v.DirectorName = v.DirectorName[:0]
	v.Alias = v.Alias[:0]
	v.SupportDefinition = v.SupportDefinition[:0]
	v.ProducerName = v.ProducerName[:0]
	v.UploaderName = v.UploaderName[:0]
	v.AlbumId = ""
	v.Description = ""
	v.Language = ""
	v.Title = ""
	v.PlayUrl = ""
	v.SubTitle = ""
	v.Id = ""
	v.Area = ""
	v.VCoverUrl = ""
	v.OupgcType = ""
	v.CoverUrl = ""
	v.Operation = ""
	v.ExtendInfo = ""
	v.ReleaseTime = 0
	v.ChargeType = 0
	v.LikeCount = 0
	v.Duration = 0
	v.IsExclusive = 0
	v.ValidEndTime = 0
	v.ForwardCount = 0
	v.CommonCateId = 0
	v.ValidStartTime = 0
	v.CollectCount = 0
	v.CommentCount = 0
	v.PlayCount = 0
	v.OrderIndex = 0
	v.ContentScore = 0
	v.SuggestMaxPrice = 0
	v.SuggestMinPrice = 0
	v.CostPrice = 0
	v.AuditionSecond = 0
	poolRawSingleVideo.Put(v)
}
