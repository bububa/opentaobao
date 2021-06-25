package category

// Brand 
type Brand struct {

    // 对应属性的 pid:vid 串中的vid
    Vid   int64 `json:"vid,omitempty"`

    // vid的值
    Name   string `json:"name,omitempty"`

    // 品牌的属性id
    Pid   int64 `json:"pid,omitempty"`

    // 品牌的属性名
    PropName   string `json:"prop_name,omitempty"`

}
