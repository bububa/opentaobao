package campus

import (
	"sync"
)

// CalendarTemplateDto 结构体
type CalendarTemplateDto struct {
	// 时期集合
	DateList []string `json:"date_list,omitempty" xml:"date_list>string,omitempty"`
	// 日期名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 日期ID
	TemplateId int64 `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

var poolCalendarTemplateDto = sync.Pool{
	New: func() any {
		return new(CalendarTemplateDto)
	},
}

// GetCalendarTemplateDto() 从对象池中获取CalendarTemplateDto
func GetCalendarTemplateDto() *CalendarTemplateDto {
	return poolCalendarTemplateDto.Get().(*CalendarTemplateDto)
}

// ReleaseCalendarTemplateDto 释放CalendarTemplateDto
func ReleaseCalendarTemplateDto(v *CalendarTemplateDto) {
	v.DateList = v.DateList[:0]
	v.Name = ""
	v.TemplateId = 0
	poolCalendarTemplateDto.Put(v)
}
