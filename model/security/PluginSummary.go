package security

// PluginSummary 
/* model for simplify = false
type PluginSummary struct {

    // 插件个数
    
    PluginCount   int64 `json:"plugin_count,omitempty"`
    

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// PluginSummary 
type PluginSummary struct {

    // 插件个数
    PluginCount   int64 `json:"plugin_count,omitempty"`

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    Status   int64 `json:"status,omitempty"`

}
