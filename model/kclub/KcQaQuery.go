package kclub

// KcQaQuery 
/* model for simplify = false
type KcQaQuery struct {

    // 租户id
    
    TenantId   int64 `json:"tenant_id,omitempty"`
    

    // context列表
    
    ContextList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"context_list,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 问题类型列表
    
    QuestionTypes  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"question_types,omitempty"`
    

    // context
    
    Context   int64 `json:"context,omitempty"`
    

    // 页大小
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 状态列表
    
    StatusList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"status_list,omitempty"`
    

    // 类目id
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 问题类型
    
    QuestionType   int64 `json:"question_type,omitempty"`
    

    // 排序对象
    
    SorterConfig  *struct {
        SorterConfig  *SorterConfig `json:"sorter_config,omitempty"`
    } `json:"sorter_config,omitempty"`
    

    // 视角过滤
    
    Views  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"views,omitempty"`
    

}
*/

// KcQaQuery 
type KcQaQuery struct {

    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty"`

    // context列表
    ContextList   []int64 `json:"context_list,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 问题类型列表
    QuestionTypes   []int64 `json:"question_types,omitempty"`

    // context
    Context   int64 `json:"context,omitempty"`

    // 页大小
    PageSize   int64 `json:"page_size,omitempty"`

    // 状态列表
    StatusList   []int64 `json:"status_list,omitempty"`

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty"`

    // 排序对象
    SorterConfig   *SorterConfig `json:"sorter_config,omitempty"`

    // 视角过滤
    Views   []int64 `json:"views,omitempty"`

}
