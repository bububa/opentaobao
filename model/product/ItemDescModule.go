package product

// ItemDescModule 
/* model for simplify = false
type ItemDescModule struct {

    // 模块id ,(注意:用户自定义模块不用填此项。)
    
    ModuleId   int64 `json:"module_id,omitempty"`
    

    // 模块名称
    
    ModuleName   string `json:"module_name,omitempty"`
    

    // cat_mod:表示是运营设置的类目维度模块，usr_mod表示的是商家自定义模块。
    
    Type   string `json:"type,omitempty"`
    

    // 描述具体内容
    
    Content   string `json:"content,omitempty"`
    

    // 一个List<String>的Json串，里面是模块引导提示，不超过3条（isv提交时可忽略不传）
    
    Intros   string `json:"intros,omitempty"`
    

    // 淘宝图片空间的地址链接，示例模板，最多不超过三个（isv提交时可忽略不传）
    
    TplUrls   string `json:"tpl_urls,omitempty"`
    

    // 是否必填 （isv提交时可忽略不传）
    
    Required   bool `json:"required,omitempty"`
    

}
*/

// ItemDescModule 
type ItemDescModule struct {

    // 模块id ,(注意:用户自定义模块不用填此项。)
    ModuleId   int64 `json:"module_id,omitempty"`

    // 模块名称
    ModuleName   string `json:"module_name,omitempty"`

    // cat_mod:表示是运营设置的类目维度模块，usr_mod表示的是商家自定义模块。
    Type   string `json:"type,omitempty"`

    // 描述具体内容
    Content   string `json:"content,omitempty"`

    // 一个List<String>的Json串，里面是模块引导提示，不超过3条（isv提交时可忽略不传）
    Intros   string `json:"intros,omitempty"`

    // 淘宝图片空间的地址链接，示例模板，最多不超过三个（isv提交时可忽略不传）
    TplUrls   string `json:"tpl_urls,omitempty"`

    // 是否必填 （isv提交时可忽略不传）
    Required   bool `json:"required,omitempty"`

}
