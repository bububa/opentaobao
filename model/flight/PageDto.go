package flight

import (
	"sync"
)

// PageDto 结构体
type PageDto struct {
	// 结果集
	DataList []AlitripAgentCoordinateListT `json:"data_list,omitempty" xml:"data_list>alitrip_agent_coordinate_list_t,omitempty"`
	// 数据集合
	Data []AlitripAgentFlightIntentionListT `json:"data,omitempty" xml:"data>alitrip_agent_flight_intention_list_t,omitempty"`
	// 错误码:000:系统异常, 001:请求参数不合法, 002:权限不足, 003:操作失败, 004:流量管控
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 成功标示
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageDto = sync.Pool{
	New: func() any {
		return new(PageDto)
	},
}

// GetPageDto() 从对象池中获取PageDto
func GetPageDto() *PageDto {
	return poolPageDto.Get().(*PageDto)
}

// ReleasePageDto 释放PageDto
func ReleasePageDto(v *PageDto) {
	v.DataList = v.DataList[:0]
	v.Data = v.Data[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Total = 0
	v.Success = false
	poolPageDto.Put(v)
}
