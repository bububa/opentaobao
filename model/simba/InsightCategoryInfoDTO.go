package simba

// InsightCategoryInfoDTO 
/* model for simplify = false
type InsightCategoryInfoDTO struct {

    // 类目Id
    
    CatId   int64 `json:"cat_id,omitempty"`
    

    // 类目名称
    
    CatName   string `json:"cat_name,omitempty"`
    

    // 类目的级别
    
    CatLevel   int64 `json:"cat_level,omitempty"`
    

    // 父类目Id
    
    ParentCatId   int64 `json:"parent_cat_id,omitempty"`
    

    // 表示该类目的所有父级类目，按层级顺序排列，层级越高的在前面，不同的层级之间用空格分隔
    
    CatPathId   string `json:"cat_path_id,omitempty"`
    

    // 表示类目的所有父级类目的名称，不同层级之间用（ascii码为2的字符）分隔开
    
    CatPathName   string `json:"cat_path_name,omitempty"`
    

    // 表示类目信息上次同步的时间
    
    LastSyncTime   string `json:"last_sync_time,omitempty"`
    

}
*/

// InsightCategoryInfoDTO 
type InsightCategoryInfoDTO struct {

    // 类目Id
    CatId   int64 `json:"cat_id,omitempty"`

    // 类目名称
    CatName   string `json:"cat_name,omitempty"`

    // 类目的级别
    CatLevel   int64 `json:"cat_level,omitempty"`

    // 父类目Id
    ParentCatId   int64 `json:"parent_cat_id,omitempty"`

    // 表示该类目的所有父级类目，按层级顺序排列，层级越高的在前面，不同的层级之间用空格分隔
    CatPathId   string `json:"cat_path_id,omitempty"`

    // 表示类目的所有父级类目的名称，不同层级之间用（ascii码为2的字符）分隔开
    CatPathName   string `json:"cat_path_name,omitempty"`

    // 表示类目信息上次同步的时间
    LastSyncTime   string `json:"last_sync_time,omitempty"`

}
