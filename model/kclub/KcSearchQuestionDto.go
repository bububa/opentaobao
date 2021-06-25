package kclub

// KcSearchQuestionDto 
type KcSearchQuestionDto struct {

    // 生效时间
    StartDate   string `json:"start_date,omitempty"`

    // 类目id
    CatId   int64 `json:"cat_id,omitempty"`

    // 额外字段
    Ext   string `json:"ext,omitempty"`

    // 类目名称
    CatName   string `json:"cat_name,omitempty"`

    // 类目路径名称
    CatPathName   string `json:"cat_path_name,omitempty"`

    // 相似问题id
    SimilarId   int64 `json:"similar_id,omitempty"`

    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty"`

    // 编辑人
    ModifiedUserName   string `json:"modified_user_name,omitempty"`

    // 创建人
    CreateUserName   string `json:"create_user_name,omitempty"`

    // 编辑人id
    ModifiedUserId   int64 `json:"modified_user_id,omitempty"`

    // 创建人id
    CreateUserId   int64 `json:"create_user_id,omitempty"`

    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 分词
    TitleSegment   string `json:"title_segment,omitempty"`

    // 标题
    Title   string `json:"title,omitempty"`

    // 编辑时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 关联实体code
    EntityCode   string `json:"entity_code,omitempty"`

    // 父类目列表
    CatIdPathList   []Number `json:"cat_id_path_list,omitempty"`

    // 唯一编码
    Uuid   string `json:"uuid,omitempty"`

    // 得分
    Score   string `json:"score,omitempty"`

    // 相似问题标题
    SimilarTitle   string `json:"similar_title,omitempty"`

    // 索引id
    Id   int64 `json:"id,omitempty"`

    // 索引类型
    Type   int64 `json:"type,omitempty"`

    // 问题id
    QuestionId   int64 `json:"question_id,omitempty"`

    // 是否相似标题
    IsSimilarTitle   int64 `json:"is_similar_title,omitempty"`

    // 是否草稿版本
    Source   int64 `json:"source,omitempty"`

    // 版本
    Version   int64 `json:"version,omitempty"`

    // 失效时间
    EndDate   string `json:"end_date,omitempty"`

    // 十天访问pv
    TenDayPv   int64 `json:"ten_day_pv,omitempty"`

    // 7天访问pv
    SevenDayPv   int64 `json:"seven_day_pv,omitempty"`

    // 5天访问pv
    FiveDayPv   int64 `json:"five_day_pv,omitempty"`

    // 3天访问pv
    ThreeDayPv   int64 `json:"three_day_pv,omitempty"`

    // 1天访问pv
    OneDayPv   int64 `json:"one_day_pv,omitempty"`

}