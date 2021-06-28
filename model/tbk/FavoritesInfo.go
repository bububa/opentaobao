package tbk

// FavoritesInfo 
type FavoritesInfo struct {

    // 选品库总数量
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    

    // 选品库详细信息
    
    FavoritesList   []FavoritesDetail `json:"favorites_list,omitempty" xml:"favorites_list,omitempty"`
    

}
