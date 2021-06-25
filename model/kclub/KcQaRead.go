package kclub

// KcQaRead 
type KcQaRead struct {

    // 子知识
    ChildQas   []KcQaRead `json:"child_qas,omitempty"`

    // 问题答案
    Solutions   []KcQaSolution `json:"solutions,omitempty"`

    // 问题context
    Context   int64 `json:"context,omitempty"`

    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty"`

    // 问题关联实体code
    EntityCode   string `json:"entity_code,omitempty"`

    // 问题父类目列表
    ParentCats   []Number `json:"parent_cats,omitempty"`

    // 问题类目路径
    CatPath   string `json:"cat_path,omitempty"`

    // 问题类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 问题租户id
    TenantId   int64 `json:"tenant_id,omitempty"`

    // 问题状态
    Status   int64 `json:"status,omitempty"`

    // 问题标题
    Title   string `json:"title,omitempty"`

    // 问题编辑时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 问题创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 问题id
    Id   int64 `json:"id,omitempty"`

}
