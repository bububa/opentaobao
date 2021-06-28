package kclub

// KcQaRead 
type KcQaRead struct {

    // 子知识
    
    ChildQas   []KcQaRead `json:"child_qas,omitempty" xml:"child_qas,omitempty"`
    

    // 问题答案
    
    Solutions   []KcQaSolution `json:"solutions,omitempty" xml:"solutions,omitempty"`
    

    // 问题context
    
    Context   int64 `json:"context,omitempty" xml:"context,omitempty"`
    

    // 问题类型
    
    QuestionType   int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
    

    // 问题关联实体code
    
    EntityCode   string `json:"entity_code,omitempty" xml:"entity_code,omitempty"`
    

    // 问题父类目列表
    
    ParentCats   []int64 `json:"parent_cats,omitempty" xml:"parent_cats>int64,omitempty"`
    

    // 问题类目路径
    
    CatPath   string `json:"cat_path,omitempty" xml:"cat_path,omitempty"`
    

    // 问题类目id
    
    CatId   int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
    

    // 问题租户id
    
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    

    // 问题状态
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 问题标题
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 问题编辑时间
    
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    

    // 问题创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    

    // 问题id
    
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    

}
