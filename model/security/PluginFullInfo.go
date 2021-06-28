package security

// PluginFullInfo 
/* model for simplify = false
type PluginFullInfo struct {

    // 插件个数
    
    PluginCount   int64 `json:"plugin_count,omitempty"`
    

    // 插件详细信息
    
    PluginDetails  struct {
        PluginDetail  []PluginDetail `json:"plugin_detail,omitempty"`
    } `json:"plugin_details,omitempty"`
    

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    
    Status   int64 `json:"status,omitempty"`
    

}
*/

// PluginFullInfo 
type PluginFullInfo struct {

    // 插件个数
    PluginCount   int64 `json:"plugin_count,omitempty"`

    // 插件详细信息
    PluginDetails   []PluginDetail `json:"plugin_details,omitempty"`

    // 子任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
    Status   int64 `json:"status,omitempty"`

}
