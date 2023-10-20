package simba

import (
	"sync"
)

// SidVo 结构体
type SidVo struct {
	// 当前用户是否有创意本地上传试用功能
	CreativeImgUpload bool `json:"creative_img_upload,omitempty" xml:"creative_img_upload,omitempty"`
	// 当前用户是否新版直通车用户
	IsNewBPUser bool `json:"is_new_b_p_user,omitempty" xml:"is_new_b_p_user,omitempty"`
}

var poolSidVo = sync.Pool{
	New: func() any {
		return new(SidVo)
	},
}

// GetSidVo() 从对象池中获取SidVo
func GetSidVo() *SidVo {
	return poolSidVo.Get().(*SidVo)
}

// ReleaseSidVo 释放SidVo
func ReleaseSidVo(v *SidVo) {
	v.CreativeImgUpload = false
	v.IsNewBPUser = false
	poolSidVo.Put(v)
}
