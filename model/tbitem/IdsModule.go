package tbitem

import (
	"sync"
)

// IdsModule 结构体
type IdsModule struct {
	// 宝贝描述规范化模块名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 宝贝描述规范化模块id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 0为自动打标；&lt;br/&gt;1为人工打标；
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolIdsModule = sync.Pool{
	New: func() any {
		return new(IdsModule)
	},
}

// GetIdsModule() 从对象池中获取IdsModule
func GetIdsModule() *IdsModule {
	return poolIdsModule.Get().(*IdsModule)
}

// ReleaseIdsModule 释放IdsModule
func ReleaseIdsModule(v *IdsModule) {
	v.Name = ""
	v.Id = 0
	v.Type = 0
	poolIdsModule.Put(v)
}
