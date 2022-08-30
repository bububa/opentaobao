package media

// TopPictureDto 结构体
type TopPictureDto struct {
	// 图片名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 审核状态。"AUDITING":审核中, FROZEN:被冻结 ,"NORMAL":正常"
	BizStatus string `json:"biz_status,omitempty" xml:"biz_status,omitempty"`
	// "NORMAL":正常，"RECYCLED":回收站中
	DeletedStatus string `json:"deleted_status,omitempty" xml:"deleted_status,omitempty"`
	// 图片完整链接，直接通过这个链接可以访问图片
	FullUrl string `json:"full_url,omitempty" xml:"full_url,omitempty"`
	// 分辨率，格式:长x宽，如450x150
	Pixel string `json:"pixel,omitempty" xml:"pixel,omitempty"`
	// 图片扩展名
	ExtensionName string `json:"extension_name,omitempty" xml:"extension_name,omitempty"`
	// 缩略图地址
	CompressPictureUrl string `json:"compress_picture_url,omitempty" xml:"compress_picture_url,omitempty"`
	// 图片文件id
	FileId int64 `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 所属文件夹id
	FolderId int64 `json:"folder_id,omitempty" xml:"folder_id,omitempty"`
	// 图片的大小，单位为字节
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}
