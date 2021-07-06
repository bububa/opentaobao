package flight

// FlightChangeDataQueryOption 结构体
type FlightChangeDataQueryOption struct {
	// 航变信息产生时间截至,格式yyyy-MM-dd
	EndFlightChangeTimeStr string `json:"end_flight_change_time_str,omitempty" xml:"end_flight_change_time_str,omitempty"`
	// 原到达机场三字代码
	OldArrAirport string `json:"old_arr_airport,omitempty" xml:"old_arr_airport,omitempty"`
	// 原航变旅行日期开始,格式yyyy-MM-dd
	BeginOldDepTimeStr string `json:"begin_old_dep_time_str,omitempty" xml:"begin_old_dep_time_str,omitempty"`
	// 原航变旅行日期截至,格式yyyy-MM-dd
	EndOldDepTimeStr string `json:"end_old_dep_time_str,omitempty" xml:"end_old_dep_time_str,omitempty"`
	// 原出发机场三字代码
	OldDepAirport string `json:"old_dep_airport,omitempty" xml:"old_dep_airport,omitempty"`
	// 航变信息产生时间开始,格式yyyy-MM-dd
	BeginFlightChangeTimeStr string `json:"begin_flight_change_time_str,omitempty" xml:"begin_flight_change_time_str,omitempty"`
	// 原航班号
	OldFltNum string `json:"old_flt_num,omitempty" xml:"old_flt_num,omitempty"`
	// 是否只查已确认的航变,1:是 2:否
	IsConfirmed int64 `json:"is_confirmed,omitempty" xml:"is_confirmed,omitempty"`
	// 是否只查自己订单航变,1:是 2:否
	IsGetSelfOnly int64 `json:"is_get_self_only,omitempty" xml:"is_get_self_only,omitempty"`
	// 排序,1:航变时间降序（默认） 2:航变时间升序 3:离港时间降序 4:离港时间升序
	Qsort int64 `json:"qsort,omitempty" xml:"qsort,omitempty"`
	// 第几页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
