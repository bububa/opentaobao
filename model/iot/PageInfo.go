package iot

// PageInfo 
/* model for simplify = false
type PageInfo struct {

    // 分页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 分页页码
    
    PageNum   int64 `json:"page_num,omitempty"`
    

    // 总条数
    
    Total   int64 `json:"total,omitempty"`
    

    // 数据集
    
    List  struct {
        BusinessRecipeOpenDto  []BusinessRecipeOpenDto `json:"business_recipe_open_dto,omitempty"`
    } `json:"list,omitempty"`
    

}
*/

// PageInfo 
type PageInfo struct {

    // 分页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 分页页码
    PageNum   int64 `json:"page_num,omitempty"`

    // 总条数
    Total   int64 `json:"total,omitempty"`

    // 数据集
    List   []BusinessRecipeOpenDto `json:"list,omitempty"`

}
