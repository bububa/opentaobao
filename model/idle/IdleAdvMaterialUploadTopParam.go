package idle

import (
	"sync"
)

// IdleAdvMaterialUploadTopParam 结构体
type IdleAdvMaterialUploadTopParam struct {
	// 素材上传详细信息参数
	UploadDetails []IdleAdvMaterialUploadDetailTopParam `json:"upload_details,omitempty" xml:"upload_details>idle_adv_material_upload_detail_top_param,omitempty"`
}

var poolIdleAdvMaterialUploadTopParam = sync.Pool{
	New: func() any {
		return new(IdleAdvMaterialUploadTopParam)
	},
}

// GetIdleAdvMaterialUploadTopParam() 从对象池中获取IdleAdvMaterialUploadTopParam
func GetIdleAdvMaterialUploadTopParam() *IdleAdvMaterialUploadTopParam {
	return poolIdleAdvMaterialUploadTopParam.Get().(*IdleAdvMaterialUploadTopParam)
}

// ReleaseIdleAdvMaterialUploadTopParam 释放IdleAdvMaterialUploadTopParam
func ReleaseIdleAdvMaterialUploadTopParam(v *IdleAdvMaterialUploadTopParam) {
	v.UploadDetails = v.UploadDetails[:0]
	poolIdleAdvMaterialUploadTopParam.Put(v)
}
