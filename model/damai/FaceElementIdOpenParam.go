package damai

import (
	"sync"
)

// FaceElementIdOpenParam 结构体
type FaceElementIdOpenParam struct {
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
}

var poolFaceElementIdOpenParam = sync.Pool{
	New: func() any {
		return new(FaceElementIdOpenParam)
	},
}

// GetFaceElementIdOpenParam() 从对象池中获取FaceElementIdOpenParam
func GetFaceElementIdOpenParam() *FaceElementIdOpenParam {
	return poolFaceElementIdOpenParam.Get().(*FaceElementIdOpenParam)
}

// ReleaseFaceElementIdOpenParam 释放FaceElementIdOpenParam
func ReleaseFaceElementIdOpenParam(v *FaceElementIdOpenParam) {
	v.SupplierSecret = ""
	v.FaceId = 0
	v.SystemId = 0
	poolFaceElementIdOpenParam.Put(v)
}
