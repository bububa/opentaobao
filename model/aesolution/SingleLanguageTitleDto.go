package aesolution

import (
	"sync"
)

// SingleLanguageTitleDto 结构体
type SingleLanguageTitleDto struct {
	// Support: en(English) ru(Russian) es(Spanish) fr(French) it(Italian) tr(Turkish) pt(Portuguese) de(German) nl(Dutch) in(Indonesian) ar(Arabic) ja(Japanese) ko(Korean) th(Thai) vi(Vietnamese) iw(Hebrew).Must contains the original locale.
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// subject, maximum 218 characters.
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
}

var poolSingleLanguageTitleDto = sync.Pool{
	New: func() any {
		return new(SingleLanguageTitleDto)
	},
}

// GetSingleLanguageTitleDto() 从对象池中获取SingleLanguageTitleDto
func GetSingleLanguageTitleDto() *SingleLanguageTitleDto {
	return poolSingleLanguageTitleDto.Get().(*SingleLanguageTitleDto)
}

// ReleaseSingleLanguageTitleDto 释放SingleLanguageTitleDto
func ReleaseSingleLanguageTitleDto(v *SingleLanguageTitleDto) {
	v.Language = ""
	v.Subject = ""
	poolSingleLanguageTitleDto.Put(v)
}
