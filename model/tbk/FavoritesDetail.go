package tbk

// FavoritesDetail 结构体
type FavoritesDetail struct {
	// 选品库标题
	Favoritestitle string `json:"favorites_title,omitempty" xml:"favorites_title,omitempty"`
	// 选品库id
	Favoritesid int64 `json:"favorites_id,omitempty" xml:"favorites_id,omitempty"`
}
