package travel

// ItemScenicInfo 
/* model for simplify = false
type ItemScenicInfo struct {

    // 结构化景点信息 景点结构化信息和文本描述二选一
    
    ScenicList  struct {
        ItemScenic  []ItemScenic `json:"item_scenic,omitempty"`
    } `json:"scenic_list,omitempty"`
    

}
*/

// ItemScenicInfo 
type ItemScenicInfo struct {

    // 结构化景点信息 景点结构化信息和文本描述二选一
    ScenicList   []ItemScenic `json:"scenic_list,omitempty"`

}
