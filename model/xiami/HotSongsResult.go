package xiami

// HotSongsResult 
/* model for simplify = false
type HotSongsResult struct {

    // 歌曲列表
    
    Songs  struct {
        StandardSong  []StandardSong `json:"standard_song,omitempty"`
    } `json:"songs,omitempty"`
    

}
*/

// HotSongsResult 
type HotSongsResult struct {

    // 歌曲列表
    Songs   []StandardSong `json:"songs,omitempty"`

}
