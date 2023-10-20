package fenxiao

import (
	"sync"
)

// TopMemoDto 结构体
type TopMemoDto struct {
	// 附件
	Attachments []TopMemoAttachment `json:"attachments,omitempty" xml:"attachments>top_memo_attachment,omitempty"`
	// 子订单备注内容
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 操作者昵称
	OperateUserNick string `json:"operate_user_nick,omitempty" xml:"operate_user_nick,omitempty"`
}

var poolTopMemoDto = sync.Pool{
	New: func() any {
		return new(TopMemoDto)
	},
}

// GetTopMemoDto() 从对象池中获取TopMemoDto
func GetTopMemoDto() *TopMemoDto {
	return poolTopMemoDto.Get().(*TopMemoDto)
}

// ReleaseTopMemoDto 释放TopMemoDto
func ReleaseTopMemoDto(v *TopMemoDto) {
	v.Attachments = v.Attachments[:0]
	v.Remark = ""
	v.OperateUserNick = ""
	poolTopMemoDto.Put(v)
}
