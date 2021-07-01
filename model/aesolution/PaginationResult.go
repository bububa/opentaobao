package aesolution

// PaginationResult 
type PaginationResult struct {
    // error massage
    ErrorMessage   string `json:"error_message,omitempty" xml:"error_message,omitempty"`
    // total count(SC order is not include the result）
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // target list
    TargetList   []OrderDto `json:"target_list,omitempty" xml:"target_list>order_dto,omitempty"`
    // the number of each page
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // error code
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // current page
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // total page
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // success or not
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // timeStamp
    TimeStamp   string `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
}
