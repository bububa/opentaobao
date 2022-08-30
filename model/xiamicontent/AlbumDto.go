package xiamicontent

// AlbumDto 结构体
type AlbumDto struct {
	// 专辑封面
	AlbumLogo string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
	// 专辑名
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
	// 副标题
	SubName string `json:"sub_name,omitempty" xml:"sub_name,omitempty"`
	// 发布时间
	GmtPublish string `json:"gmt_publish,omitempty" xml:"gmt_publish,omitempty"`
	// 语种
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 厂牌id
	CompanyId int64 `json:"company_id,omitempty" xml:"company_id,omitempty"`
	// 专辑状态
	AlbumStatus int64 `json:"album_status,omitempty" xml:"album_status,omitempty"`
	// 专辑id
	AlbumId int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
}
