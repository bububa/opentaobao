package media

// File 结构体
type File struct {
	// 文件在多媒体平台的编号
	FileId int64 `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 该文件所属目录的目录编号
	DirId int64 `json:"dir_id,omitempty" xml:"dir_id,omitempty"`
	// 文件在多媒体平台的文件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
	FilePath string `json:"file_path,omitempty" xml:"file_path,omitempty"`
	// 文件的大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 申请cdn资源的分配的userId
	CdnUserId int64 `json:"cdn_user_id,omitempty" xml:"cdn_user_id,omitempty"`
	// 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 文件是否删除
	Deleted string `json:"deleted,omitempty" xml:"deleted,omitempty"`
	// 文件内容修改时间
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 图片像素，如果非图片，该值为空
	PicturePix string `json:"picture_pix,omitempty" xml:"picture_pix,omitempty"`
}
