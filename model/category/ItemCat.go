package category

// ItemCat 
/* model for simplify = false
type ItemCat struct {

    // 商品所属类目ID
    
    Cid   int64 `json:"cid,omitempty"`
    

    // 父类目ID=0时，代表的是一级的类目
    
    ParentCid   int64 `json:"parent_cid,omitempty"`
    

    // 类目名称
    
    Name   string `json:"name,omitempty"`
    

    // 该类目是否为父类目(即：该类目是否还有子类目)
    
    IsParent   bool `json:"is_parent,omitempty"`
    

    // 状态。可选值:normal(正常),deleted(删除)
    
    Status   string `json:"status,omitempty"`
    

    // 排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数
    
    SortOrder   int64 `json:"sort_order,omitempty"`
    

    // Feature对象列表<br/>目前已有的属性：<br/>若Attr_key为 udsaleprop，attr_value为1 则允许卖家在改类目新增自定义销售属性,不然为不允许
    
    Features  struct {
        Feature  []Feature `json:"feature,omitempty"`
    } `json:"features,omitempty"`
    

    // 是否度量衡类目
    
    TaosirCat   bool `json:"taosir_cat,omitempty"`
    

}
*/

// ItemCat 
type ItemCat struct {

    // 商品所属类目ID
    Cid   int64 `json:"cid,omitempty"`

    // 父类目ID=0时，代表的是一级的类目
    ParentCid   int64 `json:"parent_cid,omitempty"`

    // 类目名称
    Name   string `json:"name,omitempty"`

    // 该类目是否为父类目(即：该类目是否还有子类目)
    IsParent   bool `json:"is_parent,omitempty"`

    // 状态。可选值:normal(正常),deleted(删除)
    Status   string `json:"status,omitempty"`

    // 排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数
    SortOrder   int64 `json:"sort_order,omitempty"`

    // Feature对象列表<br/>目前已有的属性：<br/>若Attr_key为 udsaleprop，attr_value为1 则允许卖家在改类目新增自定义销售属性,不然为不允许
    Features   []Feature `json:"features,omitempty"`

    // 是否度量衡类目
    TaosirCat   bool `json:"taosir_cat,omitempty"`

}
