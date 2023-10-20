package servicecenter

import (
	"sync"
)

// CreditInfoTopDto 结构体
type CreditInfoTopDto struct {
	// 身份证
	IdentityNo string `json:"identity_no,omitempty" xml:"identity_no,omitempty"`
	// 名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 被拒原因，只支持传入1,2,3,4.其中1代表综合评分不足；2代表黑名单客群；3代表信用不良；4代表其他
	RejectMsg string `json:"reject_msg,omitempty" xml:"reject_msg,omitempty"`
	// 唯一id
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 手机号
	Mobile int64 `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 额度，单位分，假设离线通过的不用给额度
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 0代表已经完成，1代表还需要进一步评估处理
	Flag int64 `json:"flag,omitempty" xml:"flag,omitempty"`
	// 是否通过
	Pass bool `json:"pass,omitempty" xml:"pass,omitempty"`
}

var poolCreditInfoTopDto = sync.Pool{
	New: func() any {
		return new(CreditInfoTopDto)
	},
}

// GetCreditInfoTopDto() 从对象池中获取CreditInfoTopDto
func GetCreditInfoTopDto() *CreditInfoTopDto {
	return poolCreditInfoTopDto.Get().(*CreditInfoTopDto)
}

// ReleaseCreditInfoTopDto 释放CreditInfoTopDto
func ReleaseCreditInfoTopDto(v *CreditInfoTopDto) {
	v.IdentityNo = ""
	v.Name = ""
	v.RejectMsg = ""
	v.Uuid = ""
	v.Mobile = 0
	v.Amount = 0
	v.Flag = 0
	v.Pass = false
	poolCreditInfoTopDto.Put(v)
}
