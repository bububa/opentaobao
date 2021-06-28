package xiami

// TagGenreSongresult 
/* model for simplify = false
type TagGenreSongresult struct {

    // 总数
    
    Total   int64 `json:"total,omitempty"`
    

    // 是否还有
    
    More   bool `json:"more,omitempty"`
    

    // 歌曲列表
    
    Songs  struct {
        Songs  []Songs `json:"songs,omitempty"`
    } `json:"songs,omitempty"`
    

}
*/

// TagGenreSongresult 
type TagGenreSongresult struct {

    // 总数
    Total   int64 `json:"total,omitempty"`

    // 是否还有
    More   bool `json:"more,omitempty"`

    // 歌曲列表
    Songs   []Songs `json:"songs,omitempty"`

}
