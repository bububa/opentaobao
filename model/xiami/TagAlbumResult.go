package xiami

// TagAlbumResult 结构体
type TagAlbumResult struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// true：表示还可以继续翻页，false：到最后一页了
	More bool `json:"more,omitempty" xml:"more,omitempty"`
	// 风格专辑列表
	Albums []TagAlbums `json:"albums,omitempty" xml:"albums>tag_albums,omitempty"`
}
