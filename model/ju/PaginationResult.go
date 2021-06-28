package ju

// PaginationResult 
/* model for simplify = false
type PaginationResult struct {

    // 当前页码
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 扩展属性
    
    Extend  *struct {
        Extend  *Extend `json:"extend,omitempty"`
    } `json:"extend,omitempty"`
    

    // 商品数据
    
    ModelList  struct {
        Items  []Items `json:"items,omitempty"`
    } `json:"model_list,omitempty"`
    

    // 错误码
    
    MsgCode   string `json:"msg_code,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

    // 一页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 请求是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 商品总数
    
    TotalItem   int64 `json:"total_item,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 埋点信息
    
    TrackParams  *struct {
        Trackparams  *Trackparams `json:"trackparams,omitempty"`
    } `json:"track_params,omitempty"`
    

}
*/

// PaginationResult 
type PaginationResult struct {

    // 当前页码
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 扩展属性
    Extend   *Extend `json:"extend,omitempty"`

    // 商品数据
    ModelList   []Items `json:"model_list,omitempty"`

    // 错误码
    MsgCode   string `json:"msg_code,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 一页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 请求是否成功
    Success   bool `json:"success,omitempty"`

    // 商品总数
    TotalItem   int64 `json:"total_item,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 埋点信息
    TrackParams   *Trackparams `json:"track_params,omitempty"`

}
