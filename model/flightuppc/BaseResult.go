package flightuppc

import (
	"sync"
)

// BaseResult 结构体
type BaseResult struct {
	// 数据实体
	DataList []AlitripFlightBasicDataCityQueryAllData `json:"data_list,omitempty" xml:"data_list>alitrip_flight_basic_data_city_query_all_data,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 请求唯一标识
	QueryId string `json:"query_id,omitempty" xml:"query_id,omitempty"`
	// 结果码 0成功
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolBaseResult = sync.Pool{
	New: func() any {
		return new(BaseResult)
	},
}

// GetBaseResult() 从对象池中获取BaseResult
func GetBaseResult() *BaseResult {
	return poolBaseResult.Get().(*BaseResult)
}

// ReleaseBaseResult 释放BaseResult
func ReleaseBaseResult(v *BaseResult) {
	v.DataList = v.DataList[:0]
	v.ErrorMessage = ""
	v.QueryId = ""
	v.ResultCode = 0
	poolBaseResult.Put(v)
}
