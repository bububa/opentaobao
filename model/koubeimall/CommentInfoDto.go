package koubeimall

import (
	"sync"
)

// CommentInfoDto 结构体
type CommentInfoDto struct {
	// 评论详情信息list
	DetailInfoList []CommentDetailInfo `json:"detail_info_list,omitempty" xml:"detail_info_list>comment_detail_info,omitempty"`
	// 评论标签list
	CommentTagList []CommentTag `json:"comment_tag_list,omitempty" xml:"comment_tag_list>comment_tag,omitempty"`
	// 平均分
	AvgScore string `json:"avg_score,omitempty" xml:"avg_score,omitempty"`
	// 带图的评论数
	ImageCount int64 `json:"image_count,omitempty" xml:"image_count,omitempty"`
	// 总评论数
	TotalComments int64 `json:"total_comments,omitempty" xml:"total_comments,omitempty"`
}

var poolCommentInfoDto = sync.Pool{
	New: func() any {
		return new(CommentInfoDto)
	},
}

// GetCommentInfoDto() 从对象池中获取CommentInfoDto
func GetCommentInfoDto() *CommentInfoDto {
	return poolCommentInfoDto.Get().(*CommentInfoDto)
}

// ReleaseCommentInfoDto 释放CommentInfoDto
func ReleaseCommentInfoDto(v *CommentInfoDto) {
	v.DetailInfoList = v.DetailInfoList[:0]
	v.CommentTagList = v.CommentTagList[:0]
	v.AvgScore = ""
	v.ImageCount = 0
	v.TotalComments = 0
	poolCommentInfoDto.Put(v)
}
