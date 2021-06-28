package kclub

// KcSearchQuestion 
/* model for simplify = false
type KcSearchQuestion struct {

    // 问题关联的实体code
    
    EntityCode   string `json:"entity_code,omitempty"`
    

    // 问题类型
    
    QuestionType   int64 `json:"question_type,omitempty"`
    

    // 编辑人名称
    
    ModifiedUserName   string `json:"modified_user_name,omitempty"`
    

    // 创建人名称
    
    CreateUserName   string `json:"create_user_name,omitempty"`
    

    // 租户id
    
    TenantId   int64 `json:"tenant_id,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 标题分词
    
    TitleSegment   string `json:"title_segment,omitempty"`
    

    // 标题
    
    Title   string `json:"title,omitempty"`
    

    // 编辑时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 问题id
    
    Id   int64 `json:"id,omitempty"`
    

    // 问题类型：0-普通/1-引用
    
    Type   int64 `json:"type,omitempty"`
    

    // 问题id
    
    QuestionId   int64 `json:"question_id,omitempty"`
    

    // 是否相似标题
    
    IsSimilarTitle   int64 `json:"is_similar_title,omitempty"`
    

    // 来源：草稿表-0/正式表-1
    
    Source   int64 `json:"source,omitempty"`
    

    // 版本号
    
    Version   int64 `json:"version,omitempty"`
    

    // qa 失效时间
    
    EndDate   string `json:"end_date,omitempty"`
    

    // qa 生效时间
    
    StartDate   string `json:"start_date,omitempty"`
    

    // 类目id
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 额外字段
    
    Ext   string `json:"ext,omitempty"`
    

    // 类目名称
    
    CatName   string `json:"cat_name,omitempty"`
    

    // 类目路径
    
    CatPathName   string `json:"cat_path_name,omitempty"`
    

    // 相似问题主键id
    
    SimilarId   int64 `json:"similar_id,omitempty"`
    

    // catId 路径
    
    CatIdPathList  struct {
        Number  []int64 `json:"int64,omitempty"`
    } `json:"cat_id_path_list,omitempty"`
    

    // 记录唯一编码
    
    Uuid   string `json:"uuid,omitempty"`
    

    // 得分
    
    Score   string `json:"score,omitempty"`
    

    // 相似标题
    
    SimilarTitle   string `json:"similar_title,omitempty"`
    

}
*/

// KcSearchQuestion 
type KcSearchQuestion struct {

    // 问题关联的实体code
    EntityCode   string `json:"entity_code,omitempty"`

    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty"`

    // 编辑人名称
    ModifiedUserName   string `json:"modified_user_name,omitempty"`

    // 创建人名称
    CreateUserName   string `json:"create_user_name,omitempty"`

    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 标题分词
    TitleSegment   string `json:"title_segment,omitempty"`

    // 标题
    Title   string `json:"title,omitempty"`

    // 编辑时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 问题id
    Id   int64 `json:"id,omitempty"`

    // 问题类型：0-普通/1-引用
    Type   int64 `json:"type,omitempty"`

    // 问题id
    QuestionId   int64 `json:"question_id,omitempty"`

    // 是否相似标题
    IsSimilarTitle   int64 `json:"is_similar_title,omitempty"`

    // 来源：草稿表-0/正式表-1
    Source   int64 `json:"source,omitempty"`

    // 版本号
    Version   int64 `json:"version,omitempty"`

    // qa 失效时间
    EndDate   string `json:"end_date,omitempty"`

    // qa 生效时间
    StartDate   string `json:"start_date,omitempty"`

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 额外字段
    Ext   string `json:"ext,omitempty"`

    // 类目名称
    CatName   string `json:"cat_name,omitempty"`

    // 类目路径
    CatPathName   string `json:"cat_path_name,omitempty"`

    // 相似问题主键id
    SimilarId   int64 `json:"similar_id,omitempty"`

    // catId 路径
    CatIdPathList   []int64 `json:"cat_id_path_list,omitempty"`

    // 记录唯一编码
    Uuid   string `json:"uuid,omitempty"`

    // 得分
    Score   string `json:"score,omitempty"`

    // 相似标题
    SimilarTitle   string `json:"similar_title,omitempty"`

}
