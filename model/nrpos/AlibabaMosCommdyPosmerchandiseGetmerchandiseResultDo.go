package nrpos

import (
	"sync"
)

// AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo 结构体
type AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo struct {
	// 返回数据，
	Datas []MerchandiseInfoDto `json:"datas,omitempty" xml:"datas>merchandise_info_dto,omitempty"`
	// 返回头
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// null
	MappingCode string `json:"mapping_code,omitempty" xml:"mapping_code,omitempty"`
	// null
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 业务返回码，801表示商品不存在
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 业务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// http请求返回码
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo)
	},
}

// GetAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo() 从对象池中获取AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo
func GetAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo() *AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo {
	return poolAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo.Get().(*AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo)
}

// ReleaseAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo 释放AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo
func ReleaseAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo(v *AlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo) {
	v.Datas = v.Datas[:0]
	v.Headers = ""
	v.MappingCode = ""
	v.BizExtMap = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolAlibabaMosCommdyPosmerchandiseGetmerchandiseResultDo.Put(v)
}
