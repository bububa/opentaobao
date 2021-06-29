package kclub

// KcSearchQuestionQuery 
type KcSearchQuestionQuery struct {
    // 租户id
    TenantId   int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
    // 租户列表
    TenantIdList   []int64 `json:"tenant_id_list,omitempty" xml:"tenant_id_list>int64,omitempty"`
    // 搜索字段
    SearchString   string `json:"search_string,omitempty" xml:"search_string,omitempty"`
    // 问题id
    KnowledgeId   int64 `json:"knowledge_id,omitempty" xml:"knowledge_id,omitempty"`
    // 高亮显示
    Highlight   bool `json:"highlight,omitempty" xml:"highlight,omitempty"`
    // 编辑人
    EditorId   int64 `json:"editor_id,omitempty" xml:"editor_id,omitempty"`
    // 问题类型
    QuestionType   int64 `json:"question_type,omitempty" xml:"question_type,omitempty"`
    // 是否需要返回内容
    NeedContent   bool `json:"need_content,omitempty" xml:"need_content,omitempty"`
    // context列表
    ContextList   []int64 `json:"context_list,omitempty" xml:"context_list>int64,omitempty"`
    // 状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 创建者id
    CreatorId   int64 `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
    // 是否包含子类目
    IncludeSubCategorys   bool `json:"include_sub_categorys,omitempty" xml:"include_sub_categorys,omitempty"`
    // 视角列表
    Views   []int64 `json:"views,omitempty" xml:"views>int64,omitempty"`
    // 类目id列表
    CategoryIdList   []int64 `json:"category_id_list,omitempty" xml:"category_id_list>int64,omitempty"`
    // 当前页
    CurrentPage   int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
    // 是否默认展示扩展标题
    DisplayExtTitle   bool `json:"display_ext_title,omitempty" xml:"display_ext_title,omitempty"`
    // 排序规则
    SearchRule   string `json:"search_rule,omitempty" xml:"search_rule,omitempty"`
    // 排序配置
    SorterConfig   *SorterConfig `json:"sorter_config,omitempty" xml:"sorter_config,omitempty"`
}
