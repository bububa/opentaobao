package tvupadmin

// DfPageResultDTO 
type DfPageResultDTO struct {
    // code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // codeName
    CodeName   string `json:"code_name,omitempty" xml:"code_name,omitempty"`
    // detailMessage
    DetailMessage   string `json:"detail_message,omitempty" xml:"detail_message,omitempty"`
    // message
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // msgCode
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // msgInfo
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
    // pageNo
    PageNo   int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
    // pageSize
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // resultCode
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // totalCount
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
    // totalPage
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
    // value
    Values   []string `json:"values,omitempty" xml:"values>string,omitempty"`
}
