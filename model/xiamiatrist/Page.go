package xiamiatrist

// Page 
type Page struct {
    // 艺人信息列表
    Artists   []ArtistDto `json:"artists,omitempty" xml:"artists>artist_dto,omitempty"`
    // 满足条件的总数
    Count   int64 `json:"count,omitempty" xml:"count,omitempty"`
    // 分页信息
    Paging   *Paging `json:"paging,omitempty" xml:"paging,omitempty"`
}
