package logistic

// Instps 
type Instps struct {

    // 公司名称
    Name   string `json:"name,omitempty"`

    // 公司编码
    Code   string `json:"code,omitempty"`

    // 是否商家绑定的默认安装公司
    IsDefault   bool `json:"is_default,omitempty"`

}
