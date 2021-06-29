package security

// PluginDetail 
type PluginDetail struct {
    // 插件行为
    Actions   []string `json:"actions,omitempty" xml:"actions>string,omitempty"`
    // 插件开发商
    Company   string `json:"company,omitempty" xml:"company,omitempty"`
    // 插件描述
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    // 插件名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 插件类型
    Types   []string `json:"types,omitempty" xml:"types>string,omitempty"`
    // 插件位置
    Path   string `json:"path,omitempty" xml:"path,omitempty"`
}
