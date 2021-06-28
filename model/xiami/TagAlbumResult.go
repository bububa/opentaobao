package xiami

// TagAlbumResult 
/* model for simplify = false
type TagAlbumResult struct {

    // 总数
    
    Total   int64 `json:"total,omitempty"`
    

    // true：表示还可以继续翻页，false：到最后一页了
    
    More   bool `json:"more,omitempty"`
    

    // 风格专辑列表
    
    Albums  struct {
        TagAlbums  []TagAlbums `json:"tag_albums,omitempty"`
    } `json:"albums,omitempty"`
    

}
*/

// TagAlbumResult 
type TagAlbumResult struct {

    // 总数
    Total   int64 `json:"total,omitempty"`

    // true：表示还可以继续翻页，false：到最后一页了
    More   bool `json:"more,omitempty"`

    // 风格专辑列表
    Albums   []TagAlbums `json:"albums,omitempty"`

}
