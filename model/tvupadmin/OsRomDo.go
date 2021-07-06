package tvupadmin

// OsRomDo 结构体
type OsRomDo struct {
	// 下载地址
	DownloadPath string `json:"download_path,omitempty" xml:"download_path,omitempty"`
	// 文件MD5
	Downloadmd5 string `json:"downloadmd5,omitempty" xml:"downloadmd5,omitempty"`
	// 大小
	Size string `json:"size,omitempty" xml:"size,omitempty"`
	// 是否删除
	IsDelete string `json:"is_delete,omitempty" xml:"is_delete,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModify string `json:"gmt_modify,omitempty" xml:"gmt_modify,omitempty"`
	// 主键
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 关联的版本ID
	VersionId int64 `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// 基准版本ID
	BaseVersionId int64 `json:"base_version_id,omitempty" xml:"base_version_id,omitempty"`
	// 切片数
	Splitnum int64 `json:"splitnum,omitempty" xml:"splitnum,omitempty"`
}
