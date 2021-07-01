package xhotelitem

// TaobaoXhotelMultiplerateDeleteResultSet 结构体
type TaobaoXhotelMultiplerateDeleteResultSet struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 房型编码_房价编码_入住人数_连住天数
	OutRidRateplanCodeOccupancyLengthofstay string `json:"out_rid_rateplan_code_occupancy_lengthofstay,omitempty" xml:"out_rid_rateplan_code_occupancy_lengthofstay,omitempty"`
}
