package tmallnr

import (
	"sync"
)

// NrFetchCodeQueryRespDto 结构体
type NrFetchCodeQueryRespDto struct {
	// 实体信息
	NrFetchCodeDTOList []NrFetchCodeDto `json:"nr_fetch_code_d_t_o_list,omitempty" xml:"nr_fetch_code_d_t_o_list>nr_fetch_code_dto,omitempty"`
	// 返回的订单错误信息映射
	ErrMsg *Errmsg `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}

var poolNrFetchCodeQueryRespDto = sync.Pool{
	New: func() any {
		return new(NrFetchCodeQueryRespDto)
	},
}

// GetNrFetchCodeQueryRespDto() 从对象池中获取NrFetchCodeQueryRespDto
func GetNrFetchCodeQueryRespDto() *NrFetchCodeQueryRespDto {
	return poolNrFetchCodeQueryRespDto.Get().(*NrFetchCodeQueryRespDto)
}

// ReleaseNrFetchCodeQueryRespDto 释放NrFetchCodeQueryRespDto
func ReleaseNrFetchCodeQueryRespDto(v *NrFetchCodeQueryRespDto) {
	v.NrFetchCodeDTOList = v.NrFetchCodeDTOList[:0]
	v.ErrMsg = nil
	poolNrFetchCodeQueryRespDto.Put(v)
}
