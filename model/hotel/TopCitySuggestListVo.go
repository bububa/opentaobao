package hotel

// TopCitySuggestListVo 
/* model for simplify = false
type TopCitySuggestListVo struct {

    // 联想词列表
    
    SuggestItemVOList  struct {
        SuggestItemVo  []SuggestItemVo `json:"suggest_item_vo,omitempty"`
    } `json:"suggest_item_v_o_list,omitempty"`
    

}
*/

// TopCitySuggestListVo 
type TopCitySuggestListVo struct {

    // 联想词列表
    SuggestItemVOList   []SuggestItemVo `json:"suggest_item_v_o_list,omitempty"`

}
