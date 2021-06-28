package tbk

// FavoritesInfo 
/* model for simplify = false
type FavoritesInfo struct {

    // 选品库总数量
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 选品库详细信息
    
    FavoritesList  struct {
        FavoritesDetail  []FavoritesDetail `json:"favorites_detail,omitempty"`
    } `json:"favorites_list,omitempty"`
    

}
*/

// FavoritesInfo 
type FavoritesInfo struct {

    // 选品库总数量
    TotalCount   int64 `json:"total_count,omitempty"`

    // 选品库详细信息
    FavoritesList   []FavoritesDetail `json:"favorites_list,omitempty"`

}
