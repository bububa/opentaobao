package iotticket

import (
	"sync"
)

// CommentTicketTopRequest 结构体
type CommentTicketTopRequest struct {
	// 快递凭证照片
	SendProof []string `json:"send_proof,omitempty" xml:"send_proof>string,omitempty"`
	// 扩展信息
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 操作人联系方式
	OperatorPhone string `json:"operator_phone,omitempty" xml:"operator_phone,omitempty"`
	// 操作人编码
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 服务商唯一编码
	SpCode string `json:"sp_code,omitempty" xml:"sp_code,omitempty"`
	// 工单备注
	Comment string `json:"comment,omitempty" xml:"comment,omitempty"`
	// 邮寄编码
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 工单Id
	TicketId int64 `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
}

var poolCommentTicketTopRequest = sync.Pool{
	New: func() any {
		return new(CommentTicketTopRequest)
	},
}

// GetCommentTicketTopRequest() 从对象池中获取CommentTicketTopRequest
func GetCommentTicketTopRequest() *CommentTicketTopRequest {
	return poolCommentTicketTopRequest.Get().(*CommentTicketTopRequest)
}

// ReleaseCommentTicketTopRequest 释放CommentTicketTopRequest
func ReleaseCommentTicketTopRequest(v *CommentTicketTopRequest) {
	v.SendProof = v.SendProof[:0]
	v.Feature = ""
	v.OperatorPhone = ""
	v.OperatorId = ""
	v.OperatorName = ""
	v.SpCode = ""
	v.Comment = ""
	v.MailNo = ""
	v.TicketId = 0
	poolCommentTicketTopRequest.Put(v)
}
