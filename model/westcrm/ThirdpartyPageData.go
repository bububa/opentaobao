package westcrm

// ThirdpartyPageData 
type ThirdpartyPageData struct {

    // 当前页
    
    Current   int64 `json:"current,omitempty" xml:"current,omitempty"`
    

    // 活动列表
    
    List   []ActivitiesListVo `json:"list,omitempty" xml:"list,omitempty"`
    

    // 总数
    
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
    

}
