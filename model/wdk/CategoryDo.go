package wdk

// CategoryDo 
type CategoryDo struct {

    // 类目编码
    Code   string `json:"code,omitempty"`

    // 类目名称
    Name   string `json:"name,omitempty"`

    // 父类目编码（为空则认为是新增顶级类目）
    ParentCode   string `json:"parent_code,omitempty"`

    // 是否是叶子节点（叶子类目则不允许再添加子类目,非叶子类目不允许添加商品，默认false）
    Leaf   bool `json:"leaf,omitempty"`

    // 描述
    Desc   string `json:"desc,omitempty"`

    // 类目排序值
    SortOrder   int64 `json:"sort_order,omitempty"`

    // 状态（0:正常;1:禁用，禁用后已有的商品不影响，不能新建新商品）
    Status   int64 `json:"status,omitempty"`

}
