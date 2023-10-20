package nrpos

import (
	"sync"
)

// AlibabaMosCommdyOfflineGetfileurlResultDo 结构体
type AlibabaMosCommdyOfflineGetfileurlResultDo struct {
	// 返回结果合集
	Datas []OfflineFileDto `json:"datas,omitempty" xml:"datas>offline_file_dto,omitempty"`
	// 返回头
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// null
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// null
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 业务返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 业务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// http请求返回码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosCommdyOfflineGetfileurlResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosCommdyOfflineGetfileurlResultDo)
	},
}

// GetAlibabaMosCommdyOfflineGetfileurlResultDo() 从对象池中获取AlibabaMosCommdyOfflineGetfileurlResultDo
func GetAlibabaMosCommdyOfflineGetfileurlResultDo() *AlibabaMosCommdyOfflineGetfileurlResultDo {
	return poolAlibabaMosCommdyOfflineGetfileurlResultDo.Get().(*AlibabaMosCommdyOfflineGetfileurlResultDo)
}

// ReleaseAlibabaMosCommdyOfflineGetfileurlResultDo 释放AlibabaMosCommdyOfflineGetfileurlResultDo
func ReleaseAlibabaMosCommdyOfflineGetfileurlResultDo(v *AlibabaMosCommdyOfflineGetfileurlResultDo) {
	v.Datas = v.Datas[:0]
	v.Headers = ""
	v.MappingCode = ""
	v.BizExtMap = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolAlibabaMosCommdyOfflineGetfileurlResultDo.Put(v)
}
