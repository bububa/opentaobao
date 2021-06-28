package promotion

// Page 
/* model for simplify = false
type Page struct {

    // 总数
    
    Total   int64 `json:"total,omitempty"`
    

    // 每页查询数量
    
    PageSize   int64 `json:"page_size,omitempty"`
    

    // 当前页
    
    PageNo   int64 `json:"page_no,omitempty"`
    

    // 结果
    
    Datas  struct {
        AlibabaAsrDataservicePromotionruleQueryData  []AlibabaAsrDataservicePromotionruleQueryData `json:"alibaba_asr_dataservice_promotionrule_query_data,omitempty"`
    } `json:"datas,omitempty"`
    

    // 活动列表
    
    Activities  struct {
        ActivityDto  []ActivityDto `json:"activity_dto,omitempty"`
    } `json:"activities,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 当前页
    
    CurrentPage   int64 `json:"current_page,omitempty"`
    

    // 总记录数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

}
*/

// Page 
type Page struct {

    // 总数
    Total   int64 `json:"total,omitempty"`

    // 每页查询数量
    PageSize   int64 `json:"page_size,omitempty"`

    // 当前页
    PageNo   int64 `json:"page_no,omitempty"`

    // 结果
    Datas   []AlibabaAsrDataservicePromotionruleQueryData `json:"datas,omitempty"`

    // 活动列表
    Activities   []ActivityDto `json:"activities,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty"`

    // 总记录数
    TotalCount   int64 `json:"total_count,omitempty"`

}
