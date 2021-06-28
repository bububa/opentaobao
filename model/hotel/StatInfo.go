package hotel

// StatInfo 
/* model for simplify = false
type StatInfo struct {

    // 品牌类型信息
    
    BrandTypeList  struct {
        BrandType  []BrandType `json:"brand_type,omitempty"`
    } `json:"brand_type_list,omitempty"`
    

}
*/

// StatInfo 
type StatInfo struct {

    // 品牌类型信息
    BrandTypeList   []BrandType `json:"brand_type_list,omitempty"`

}
