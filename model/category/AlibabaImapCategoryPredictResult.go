package category

// AlibabaImapCategoryPredictResult 
/* model for simplify = false
type AlibabaImapCategoryPredictResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 1
    
    TopImapUnionCategoryPathDoList  struct {
        TopImapUnionCategoryPathDo  []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do,omitempty"`
    } `json:"top_imap_union_category_path_do_list,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// AlibabaImapCategoryPredictResult 
type AlibabaImapCategoryPredictResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 1
    TopImapUnionCategoryPathDoList   []TopImapUnionCategoryPathDo `json:"top_imap_union_category_path_do_list,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
