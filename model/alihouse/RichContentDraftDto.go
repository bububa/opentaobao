package alihouse

import (
	"sync"
)

// RichContentDraftDto 结构体
type RichContentDraftDto struct {
	// 标签
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 关键字
	Keywords []string `json:"keywords,omitempty" xml:"keywords>string,omitempty"`
	// 楼盘卡片映射
	ProjectCards []ProjectCardDataDto `json:"project_cards,omitempty" xml:"project_cards>project_card_data_dto,omitempty"`
	// 视频映射
	VideoMappings []VideoDataDto `json:"video_mappings,omitempty" xml:"video_mappings>video_data_dto,omitempty"`
	// 外部富文本id
	OuterRichContentId string `json:"outer_rich_content_id,omitempty" xml:"outer_rich_content_id,omitempty"`
	// 菜鸟城市id
	CityId string `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 楼盘外部id 多个,分割
	OuterIds string `json:"outer_ids,omitempty" xml:"outer_ids,omitempty"`
	// 资讯标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 资讯短标题
	ShortTitle string `json:"short_title,omitempty" xml:"short_title,omitempty"`
	// 栏目名称
	ColumnName string `json:"column_name,omitempty" xml:"column_name,omitempty"`
	// 作者
	Author string `json:"author,omitempty" xml:"author,omitempty"`
	// 发布时间 yyyy-MM-dd HH:mm:ss格式
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 封面图
	CoverImages string `json:"cover_images,omitempty" xml:"cover_images,omitempty"`
	// 摘要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 富文本内容
	SourceRichContent string `json:"source_rich_content,omitempty" xml:"source_rich_content,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 是否原创 1 原创 0 非原创
	Original int64 `json:"original,omitempty" xml:"original,omitempty"`
	// 1 测试数据 0 正常数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolRichContentDraftDto = sync.Pool{
	New: func() any {
		return new(RichContentDraftDto)
	},
}

// GetRichContentDraftDto() 从对象池中获取RichContentDraftDto
func GetRichContentDraftDto() *RichContentDraftDto {
	return poolRichContentDraftDto.Get().(*RichContentDraftDto)
}

// ReleaseRichContentDraftDto 释放RichContentDraftDto
func ReleaseRichContentDraftDto(v *RichContentDraftDto) {
	v.Tags = v.Tags[:0]
	v.Keywords = v.Keywords[:0]
	v.ProjectCards = v.ProjectCards[:0]
	v.VideoMappings = v.VideoMappings[:0]
	v.OuterRichContentId = ""
	v.CityId = ""
	v.OuterIds = ""
	v.Title = ""
	v.ShortTitle = ""
	v.ColumnName = ""
	v.Author = ""
	v.PublishTime = ""
	v.CoverImages = ""
	v.Summary = ""
	v.SourceRichContent = ""
	v.OuterStoreId = ""
	v.Original = 0
	v.IsTest = 0
	poolRichContentDraftDto.Put(v)
}
