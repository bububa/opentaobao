package kclub

// KcQaReadDto 
type KcQaReadDto struct {

    // 类目路径
    CatPath   string `json:"cat_path,omitempty"`

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 知识标题
    Title   string `json:"title,omitempty"`

    // 编辑时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 问题id
    Id   int64 `json:"id,omitempty"`

    // 子知识
    ChildQas   []KcQaReadDto `json:"child_qas,omitempty"`

    // 问题答案
    Solutions   []KcQaSolutionDto `json:"solutions,omitempty"`

    // 问题类型(原kbs context)
    Context   int64 `json:"context,omitempty"`

    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty"`

    // 关联的实体code
    EntityCode   string `json:"entity_code,omitempty"`

    // 父类目id
    ParentCats   []Number `json:"parent_cats,omitempty"`

    // qa的访问量
    QaPv   *QaPvDto `json:"qa_pv,omitempty"`

}
