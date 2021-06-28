package xiami

// RankSongsData 
/* model for simplify = false
type RankSongsData struct {

    // 歌曲列表
    
    Songs  struct {
        RankSong  []RankSong `json:"rank_song,omitempty"`
    } `json:"songs,omitempty"`
    

}
*/

// RankSongsData 
type RankSongsData struct {

    // 歌曲列表
    Songs   []RankSong `json:"songs,omitempty"`

}
