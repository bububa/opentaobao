package security

// PluginDetail 
type PluginDetail struct {

    // 插件行为
    Actions   []String `json:"actions,omitempty"`

    // 插件开发商
    Company   string `json:"company,omitempty"`

    // 插件描述
    Desc   string `json:"desc,omitempty"`

    // 插件名称
    Name   string `json:"name,omitempty"`

    // 插件类型
    Types   []String `json:"types,omitempty"`

    // 插件位置
    Path   string `json:"path,omitempty"`

}
