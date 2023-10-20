package txcs

import (
	"sync"
)

// VerificationBillResponseDto 结构体
type VerificationBillResponseDto struct {
	// 核销单号
	VerificationNo string `json:"verification_no,omitempty" xml:"verification_no,omitempty"`
	// 核销日期
	VerifyDate string `json:"verify_date,omitempty" xml:"verify_date,omitempty"`
}

var poolVerificationBillResponseDto = sync.Pool{
	New: func() any {
		return new(VerificationBillResponseDto)
	},
}

// GetVerificationBillResponseDto() 从对象池中获取VerificationBillResponseDto
func GetVerificationBillResponseDto() *VerificationBillResponseDto {
	return poolVerificationBillResponseDto.Get().(*VerificationBillResponseDto)
}

// ReleaseVerificationBillResponseDto 释放VerificationBillResponseDto
func ReleaseVerificationBillResponseDto(v *VerificationBillResponseDto) {
	v.VerificationNo = ""
	v.VerifyDate = ""
	poolVerificationBillResponseDto.Put(v)
}
