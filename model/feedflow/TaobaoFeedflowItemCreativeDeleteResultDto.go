package feedflow

import (
	"sync"
)

// TaobaoFeedflowItemCreativeDeleteResultDto 结构体
type TaobaoFeedflowItemCreativeDeleteResultDto struct {
	// 错误信息
	ErrorList []ErrorInfoDto `json:"error_list,omitempty" xml:"error_list>error_info_dto,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFeedflowItemCreativeDeleteResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCreativeDeleteResultDto)
	},
}

// GetTaobaoFeedflowItemCreativeDeleteResultDto() 从对象池中获取TaobaoFeedflowItemCreativeDeleteResultDto
func GetTaobaoFeedflowItemCreativeDeleteResultDto() *TaobaoFeedflowItemCreativeDeleteResultDto {
	return poolTaobaoFeedflowItemCreativeDeleteResultDto.Get().(*TaobaoFeedflowItemCreativeDeleteResultDto)
}

// ReleaseTaobaoFeedflowItemCreativeDeleteResultDto 释放TaobaoFeedflowItemCreativeDeleteResultDto
func ReleaseTaobaoFeedflowItemCreativeDeleteResultDto(v *TaobaoFeedflowItemCreativeDeleteResultDto) {
	v.ErrorList = v.ErrorList[:0]
	v.Message = ""
	v.Success = false
	poolTaobaoFeedflowItemCreativeDeleteResultDto.Put(v)
}
