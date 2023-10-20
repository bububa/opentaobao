package koubeimall

import (
	"sync"
)

// CommentDetailInfo 结构体
type CommentDetailInfo struct {
	// 图片列表
	CommentImageList []string `json:"comment_image_list,omitempty" xml:"comment_image_list>string,omitempty"`
	// 评论内容
	CommentContent string `json:"comment_content,omitempty" xml:"comment_content,omitempty"`
	// 评价时间
	CommentTime string `json:"comment_time,omitempty" xml:"comment_time,omitempty"`
	// 评分
	CommentScore int64 `json:"comment_score,omitempty" xml:"comment_score,omitempty"`
}

var poolCommentDetailInfo = sync.Pool{
	New: func() any {
		return new(CommentDetailInfo)
	},
}

// GetCommentDetailInfo() 从对象池中获取CommentDetailInfo
func GetCommentDetailInfo() *CommentDetailInfo {
	return poolCommentDetailInfo.Get().(*CommentDetailInfo)
}

// ReleaseCommentDetailInfo 释放CommentDetailInfo
func ReleaseCommentDetailInfo(v *CommentDetailInfo) {
	v.CommentImageList = v.CommentImageList[:0]
	v.CommentContent = ""
	v.CommentTime = ""
	v.CommentScore = 0
	poolCommentDetailInfo.Put(v)
}
