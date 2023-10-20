package xhotelitem

// TaobaoXhotelRoomtypeConflictDataResultSet 结构体
type TaobaoXhotelRoomtypeConflictDataResultSet struct {
	// 结果集
	Results []RoomTypeCheckResultDo `json:"results,omitempty" xml:"results>room_type_check_result_do,omitempty"`
	// 警告信息
	WarnMessage string `json:"warn_message,omitempty" xml:"warn_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
