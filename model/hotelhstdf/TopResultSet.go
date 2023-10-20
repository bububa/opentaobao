package hotelhstdf

import (
	"sync"
)

// TopResultSet 结构体
type TopResultSet struct {
	// 结果集合
	ModuleList []HotelRoomStaticDo `json:"module_list,omitempty" xml:"module_list>hotel_room_static_do,omitempty"`
	// 返回列表
	Results []RoomTypePo `json:"results,omitempty" xml:"results>room_type_po,omitempty"`
	// 暂不使用
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 请求失败时返回的错误信息，一般success=false时有值
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// error_code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error_msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前查询条件下的数据总数；判断是否需要继续查询时，可以使用已请求数据量对比该值，但建议对结果集合进行判空，为空即停止
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 当前页之后是否还有数据,true--还有，可以继续请求
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResultSet = sync.Pool{
	New: func() any {
		return new(TopResultSet)
	},
}

// GetTopResultSet() 从对象池中获取TopResultSet
func GetTopResultSet() *TopResultSet {
	return poolTopResultSet.Get().(*TopResultSet)
}

// ReleaseTopResultSet 释放TopResultSet
func ReleaseTopResultSet(v *TopResultSet) {
	v.ModuleList = v.ModuleList[:0]
	v.Results = v.Results[:0]
	v.ResultCode = ""
	v.ResultMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalResults = 0
	v.HasNext = false
	v.Success = false
	poolTopResultSet.Put(v)
}
