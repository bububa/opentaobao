package flight

// FlightChangeDataDo 结构体
type FlightChangeDataDo struct {
	// 航班最新到达机场三字码, 字符长度3；仅当flightChangeType=2，该参数必填
	NewArrAirport string `json:"new_arr_airport,omitempty" xml:"new_arr_airport,omitempty"`
	// 原到达机场三字代码, 字符长度3
	OldArrAirport string `json:"old_arr_airport,omitempty" xml:"old_arr_airport,omitempty"`
	// 原航变日期 ,格式:yyyy-MM-dd
	OldDepTimeStr string `json:"old_dep_time_str,omitempty" xml:"old_dep_time_str,omitempty"`
	// 原航班号（如果是二次航变，该参数为第一航变后最新的航班号。eg:第一航变 航班号从CA123变更到CA456，那第二次航变的原航班应该为CA456）
	OldFltNum string `json:"old_flt_num,omitempty" xml:"old_flt_num,omitempty"`
	// 航班最新计划起飞时间,仅当flightChangeType=2，该参数必填；录入格式:yyyy-MM-dd HH: mm
	NewDepTimeStr string `json:"new_dep_time_str,omitempty" xml:"new_dep_time_str,omitempty"`
	// 航班最新出发机场三字码, 字符长度3；仅当flightChangeType=2，该参数必填
	NewDepAirport string `json:"new_dep_airport,omitempty" xml:"new_dep_airport,omitempty"`
	// 原出发机场三字代码, 字符长度3
	OldDepAirport string `json:"old_dep_airport,omitempty" xml:"old_dep_airport,omitempty"`
	// 变更到的最新航班号,仅当flightChangeType=2,该参数必填
	NewFltNum string `json:"new_flt_num,omitempty" xml:"new_flt_num,omitempty"`
	// 最新到达时间
	NewArrTimeStr string `json:"new_arr_time_str,omitempty" xml:"new_arr_time_str,omitempty"`
	// 航变类型,1:航班取消 2:航班变更
	FlightChangeType int64 `json:"flight_change_type,omitempty" xml:"flight_change_type,omitempty"`
	// 业务类型,0:国内机票 1:国际机票
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 飞猪机票订单号，如果输入了该参数，平台只会给该指定订单发送航变，如果不输入该参数，则会处理代理商的所有订单；正常的延误航变该参数一般不需要，如果是航班保护，大部分情况该参数应该都是必填的，因为航班保护基本每个订单保护的新航班可能都不一样。
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
