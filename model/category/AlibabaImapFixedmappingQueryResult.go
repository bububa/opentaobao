package category

// AlibabaImapFixedmappingQueryResult 
/* model for simplify = false
type AlibabaImapFixedmappingQueryResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // list参数
    
    TopImapUnionCategoryPathDoList  struct {
        TopImapUnionCategoryPathDo  []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do,omitempty"`
    } `json:"top_imap_union_category_path_do_list,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// AlibabaImapFixedmappingQueryResult 
type AlibabaImapFixedmappingQueryResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // list参数
    TopImapUnionCategoryPathDoList   []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do_list,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
