package film

// CinemaPageResult 
type CinemaPageResult struct {
    // 影院集合
    Cinemas   []CinemaPageCinemaVo `json:"cinemas,omitempty" xml:"cinemas>cinema_page_cinema_vo,omitempty"`
}
