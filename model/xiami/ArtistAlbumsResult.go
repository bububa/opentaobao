package xiami

// ArtistAlbumsResult 
type ArtistAlbumsResult struct {
    // 专辑列表
    Albums   []StandardAlbum `json:"albums,omitempty" xml:"albums>standard_album,omitempty"`
    // 总数
    TotalNumber   int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
    // 上一页
    Previous   int64 `json:"previous,omitempty" xml:"previous,omitempty"`
    // 下一页
    Next   int64 `json:"next,omitempty" xml:"next,omitempty"`
}
