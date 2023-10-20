package simba

import (
	"sync"
)

// TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto 结构体
type TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto struct {
	// 返回成功/错误码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回新质量分实体信息
	Result *QScoreSplitDto `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto)
	},
}

// GetTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto() 从对象池中获取TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto
func GetTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto() *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto {
	return poolTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto.Get().(*TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto)
}

// ReleaseTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto 释放TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto
func ReleaseTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto(v *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto) {
	v.Key = ""
	v.Message = ""
	v.Result = nil
	poolTaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto.Put(v)
}
