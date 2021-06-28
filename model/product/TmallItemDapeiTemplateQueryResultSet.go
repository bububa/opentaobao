package product

// TmallItemDapeiTemplateQueryResultSet 
/* model for simplify = false
type TmallItemDapeiTemplateQueryResultSet struct {

    // firstResult
    
    Results  struct {
        DapeiDO  []DapeiDO `json:"dapei_do,omitempty"`
    } `json:"results,omitempty"`
    

    // error
    
    Error   bool `json:"error,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // totalResults
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

    // totalPage
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // pageIndex
    
    PageIndex   int64 `json:"page_index,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmallItemDapeiTemplateQueryResultSet 
type TmallItemDapeiTemplateQueryResultSet struct {

    // firstResult
    Results   []DapeiDO `json:"results,omitempty"`

    // error
    Error   bool `json:"error,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // totalResults
    TotalResults   int64 `json:"total_results,omitempty"`

    // totalPage
    TotalPage   int64 `json:"total_page,omitempty"`

    // pageIndex
    PageIndex   int64 `json:"page_index,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

}
