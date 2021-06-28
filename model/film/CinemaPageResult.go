package film

// CinemaPageResult 
/* model for simplify = false
type CinemaPageResult struct {

    // 影院集合
    
    Cinemas  struct {
        CinemaPageCinemaVo  []CinemaPageCinemaVo `json:"cinema_page_cinema_vo,omitempty"`
    } `json:"cinemas,omitempty"`
    

}
*/

// CinemaPageResult 
type CinemaPageResult struct {

    // 影院集合
    Cinemas   []CinemaPageCinemaVo `json:"cinemas,omitempty"`

}
