package media

import (
	"sync"
)

// TopCreateFolderDto 结构体
type TopCreateFolderDto struct {
	// 创建的文件夹id
	FolderId int64 `json:"folder_id,omitempty" xml:"folder_id,omitempty"`
}

var poolTopCreateFolderDto = sync.Pool{
	New: func() any {
		return new(TopCreateFolderDto)
	},
}

// GetTopCreateFolderDto() 从对象池中获取TopCreateFolderDto
func GetTopCreateFolderDto() *TopCreateFolderDto {
	return poolTopCreateFolderDto.Get().(*TopCreateFolderDto)
}

// ReleaseTopCreateFolderDto 释放TopCreateFolderDto
func ReleaseTopCreateFolderDto(v *TopCreateFolderDto) {
	v.FolderId = 0
	poolTopCreateFolderDto.Put(v)
}
