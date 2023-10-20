package simba

import (
	"sync"
)

// TaobaoSimbaKeywordsQscoreSplitGetResultDto 结构体
type TaobaoSimbaKeywordsQscoreSplitGetResultDto struct {
	// 返回成功/错误码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回新质量分实体信息
	Result *QScoreSplitDto `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoSimbaKeywordsQscoreSplitGetResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsQscoreSplitGetResultDto)
	},
}

// GetTaobaoSimbaKeywordsQscoreSplitGetResultDto() 从对象池中获取TaobaoSimbaKeywordsQscoreSplitGetResultDto
func GetTaobaoSimbaKeywordsQscoreSplitGetResultDto() *TaobaoSimbaKeywordsQscoreSplitGetResultDto {
	return poolTaobaoSimbaKeywordsQscoreSplitGetResultDto.Get().(*TaobaoSimbaKeywordsQscoreSplitGetResultDto)
}

// ReleaseTaobaoSimbaKeywordsQscoreSplitGetResultDto 释放TaobaoSimbaKeywordsQscoreSplitGetResultDto
func ReleaseTaobaoSimbaKeywordsQscoreSplitGetResultDto(v *TaobaoSimbaKeywordsQscoreSplitGetResultDto) {
	v.Key = ""
	v.Message = ""
	v.Result = nil
	poolTaobaoSimbaKeywordsQscoreSplitGetResultDto.Put(v)
}
