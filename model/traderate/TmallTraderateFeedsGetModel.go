package traderate

import (
	"sync"
)

// TmallTraderateFeedsGetModel 结构体
type TmallTraderateFeedsGetModel struct {
	// 原始评价对应的标签列表
	Tags []Tags `json:"tags,omitempty" xml:"tags>tags,omitempty"`
	// 追加评价中带有的语义标签列表
	AppendTags []Tags `json:"append_tags,omitempty" xml:"append_tags>tags,omitempty"`
	// 追加评价内容
	AppendContent string `json:"append_content,omitempty" xml:"append_content,omitempty"`
	// 评价内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 追加评价时间
	AppendTime string `json:"append_time,omitempty" xml:"append_time,omitempty"`
	// 表示评价者的昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 评价时间
	CommentTime string `json:"comment_time,omitempty" xml:"comment_time,omitempty"`
	// ouid
	Ouid string `json:"ouid,omitempty" xml:"ouid,omitempty"`
	// 追评中是否含有负向标签
	AppendHasNegtv bool `json:"append_has_negtv,omitempty" xml:"append_has_negtv,omitempty"`
	// 原始评价是否含有负向标签
	HasNegtv bool `json:"has_negtv,omitempty" xml:"has_negtv,omitempty"`
}

var poolTmallTraderateFeedsGetModel = sync.Pool{
	New: func() any {
		return new(TmallTraderateFeedsGetModel)
	},
}

// GetTmallTraderateFeedsGetModel() 从对象池中获取TmallTraderateFeedsGetModel
func GetTmallTraderateFeedsGetModel() *TmallTraderateFeedsGetModel {
	return poolTmallTraderateFeedsGetModel.Get().(*TmallTraderateFeedsGetModel)
}

// ReleaseTmallTraderateFeedsGetModel 释放TmallTraderateFeedsGetModel
func ReleaseTmallTraderateFeedsGetModel(v *TmallTraderateFeedsGetModel) {
	v.Tags = v.Tags[:0]
	v.AppendTags = v.AppendTags[:0]
	v.AppendContent = ""
	v.Content = ""
	v.AppendTime = ""
	v.UserNick = ""
	v.CommentTime = ""
	v.Ouid = ""
	v.AppendHasNegtv = false
	v.HasNegtv = false
	poolTmallTraderateFeedsGetModel.Put(v)
}
