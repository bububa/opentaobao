package promotion

import (
	"sync"
)

// CheckToolModule 结构体
type CheckToolModule struct {
	// 工具是否检测通过。
	IsPass string `json:"is_pass,omitempty" xml:"is_pass,omitempty"`
	// 工具检测结果中的错误信息。
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 工具中已经使用的元数据。
	MetaDef string `json:"meta_def,omitempty" xml:"meta_def,omitempty"`
	// 工具审核的URL，工具检测通过后，ISV需要把该URL和工具基本信息一起提交UMP工具技术审核。
	CheckUrl string `json:"check_url,omitempty" xml:"check_url,omitempty"`
}

var poolCheckToolModule = sync.Pool{
	New: func() any {
		return new(CheckToolModule)
	},
}

// GetCheckToolModule() 从对象池中获取CheckToolModule
func GetCheckToolModule() *CheckToolModule {
	return poolCheckToolModule.Get().(*CheckToolModule)
}

// ReleaseCheckToolModule 释放CheckToolModule
func ReleaseCheckToolModule(v *CheckToolModule) {
	v.IsPass = ""
	v.ErrorMessage = ""
	v.MetaDef = ""
	v.CheckUrl = ""
	poolCheckToolModule.Put(v)
}
