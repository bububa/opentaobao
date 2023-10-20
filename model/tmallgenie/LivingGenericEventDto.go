package tmallgenie

import (
	"sync"
)

// LivingGenericEventDto 结构体
type LivingGenericEventDto struct {
	// 事件内容
	Payload string `json:"payload,omitempty" xml:"payload,omitempty"`
	// 事件命名空间
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// 命名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolLivingGenericEventDto = sync.Pool{
	New: func() any {
		return new(LivingGenericEventDto)
	},
}

// GetLivingGenericEventDto() 从对象池中获取LivingGenericEventDto
func GetLivingGenericEventDto() *LivingGenericEventDto {
	return poolLivingGenericEventDto.Get().(*LivingGenericEventDto)
}

// ReleaseLivingGenericEventDto 释放LivingGenericEventDto
func ReleaseLivingGenericEventDto(v *LivingGenericEventDto) {
	v.Payload = ""
	v.Namespace = ""
	v.Name = ""
	poolLivingGenericEventDto.Put(v)
}
