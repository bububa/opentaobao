package ticket

// ScenicAndProductResult 
/* model for simplify = false
type ScenicAndProductResult struct {

    // 景点列表
    
    ScenicList  struct {
        Scenic  []Scenic `json:"scenic,omitempty"`
    } `json:"scenic_list,omitempty"`
    

}
*/

// ScenicAndProductResult 
type ScenicAndProductResult struct {

    // 景点列表
    ScenicList   []Scenic `json:"scenic_list,omitempty"`

}
