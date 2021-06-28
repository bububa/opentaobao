package wdk

// SeasonVersionParam 
/* model for simplify = false
type SeasonVersionParam struct {

    // 商家档期号
    
    OutSeasonId   string `json:"out_season_id,omitempty"`
    

    // 参与换挡的门店列表
    
    ShopIds  struct {
        String  []string `json:"string,omitempty"`
    } `json:"shop_ids,omitempty"`
    

}
*/

// SeasonVersionParam 
type SeasonVersionParam struct {

    // 商家档期号
    OutSeasonId   string `json:"out_season_id,omitempty"`

    // 参与换挡的门店列表
    ShopIds   []string `json:"shop_ids,omitempty"`

}
