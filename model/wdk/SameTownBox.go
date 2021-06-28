package wdk

// SameTownBox 
/* model for simplify = false
type SameTownBox struct {

    // 同城包裹列表
    
    SameTownPackages  struct {
        Sametownpackages  []Sametownpackages `json:"sametownpackages,omitempty"`
    } `json:"same_town_packages,omitempty"`
    

    // 是否测试 1:测试 0:非测试
    
    IsTest   string `json:"is_test,omitempty"`
    

    // 扩展信息  json格式
    
    Attribute   string `json:"attribute,omitempty"`
    

    // 箱号
    
    ContainerCode   string `json:"container_code,omitempty"`
    

    // 箱类型 NORMAL:周转箱 COLD:冷链箱 NONE:原箱
    
    ContainerType   string `json:"container_type,omitempty"`
    

}
*/

// SameTownBox 
type SameTownBox struct {

    // 同城包裹列表
    SameTownPackages   []Sametownpackages `json:"same_town_packages,omitempty"`

    // 是否测试 1:测试 0:非测试
    IsTest   string `json:"is_test,omitempty"`

    // 扩展信息  json格式
    Attribute   string `json:"attribute,omitempty"`

    // 箱号
    ContainerCode   string `json:"container_code,omitempty"`

    // 箱类型 NORMAL:周转箱 COLD:冷链箱 NONE:原箱
    ContainerType   string `json:"container_type,omitempty"`

}
