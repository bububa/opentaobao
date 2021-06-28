package scbp

// ForbiddenKeywordBatchOperationDto 
/* model for simplify = false
type ForbiddenKeywordBatchOperationDto struct {

    // 请求参数
    
    ForbiddenKeywordOperationList  struct {
        ForbiddenKeywordOperation  []ForbiddenKeywordOperation `json:"forbidden_keyword_operation,omitempty"`
    } `json:"forbidden_keyword_operation_list,omitempty"`
    

}
*/

// ForbiddenKeywordBatchOperationDto 
type ForbiddenKeywordBatchOperationDto struct {

    // 请求参数
    ForbiddenKeywordOperationList   []ForbiddenKeywordOperation `json:"forbidden_keyword_operation_list,omitempty"`

}
