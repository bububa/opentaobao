package hotel

import (
	"sync"
)

// AlitripHotelRateGetmixratelistGetResult 结构体
type AlitripHotelRateGetmixratelistGetResult struct {
	// mappingCode
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// bizExtMap
	BizExtMap *Bizextmap `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// headers
	Headers *Headers `json:"headers,omitempty" xml:"headers,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// model
	Model *GetMixRateListResult `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripHotelRateGetmixratelistGetResult = sync.Pool{
	New: func() any {
		return new(AlitripHotelRateGetmixratelistGetResult)
	},
}

// GetAlitripHotelRateGetmixratelistGetResult() 从对象池中获取AlitripHotelRateGetmixratelistGetResult
func GetAlitripHotelRateGetmixratelistGetResult() *AlitripHotelRateGetmixratelistGetResult {
	return poolAlitripHotelRateGetmixratelistGetResult.Get().(*AlitripHotelRateGetmixratelistGetResult)
}

// ReleaseAlitripHotelRateGetmixratelistGetResult 释放AlitripHotelRateGetmixratelistGetResult
func ReleaseAlitripHotelRateGetmixratelistGetResult(v *AlitripHotelRateGetmixratelistGetResult) {
	v.MappingCode = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.BizExtMap = nil
	v.Headers = nil
	v.HttpStatusCode = 0
	v.Model = nil
	v.Success = false
	poolAlitripHotelRateGetmixratelistGetResult.Put(v)
}
