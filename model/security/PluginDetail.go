package security

// PluginDetail 
/* model for simplify = false
type PluginDetail struct {

    // 插件行为
    
    Actions  struct {
        String  []string `json:"string,omitempty"`
    } `json:"actions,omitempty"`
    

    // 插件开发商
    
    Company   string `json:"company,omitempty"`
    

    // 插件描述
    
    Desc   string `json:"desc,omitempty"`
    

    // 插件名称
    
    Name   string `json:"name,omitempty"`
    

    // 插件类型
    
    Types  struct {
        String  []string `json:"string,omitempty"`
    } `json:"types,omitempty"`
    

    // 插件位置
    
    Path   string `json:"path,omitempty"`
    

}
*/

// PluginDetail 
type PluginDetail struct {

    // 插件行为
    Actions   []string `json:"actions,omitempty"`

    // 插件开发商
    Company   string `json:"company,omitempty"`

    // 插件描述
    Desc   string `json:"desc,omitempty"`

    // 插件名称
    Name   string `json:"name,omitempty"`

    // 插件类型
    Types   []string `json:"types,omitempty"`

    // 插件位置
    Path   string `json:"path,omitempty"`

}
