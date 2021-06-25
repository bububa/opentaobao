package product

// IdsModule 
type IdsModule struct {

    // 宝贝描述规范化模块id
    Id   int64 `json:"id,omitempty"`

    // 宝贝描述规范化模块名
    Name   string `json:"name,omitempty"`

    // 0为自动打标；<br/>1为人工打标；
    Type   int64 `json:"type,omitempty"`

}
