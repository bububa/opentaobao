package wdk

// Postman 
type Postman struct {

    // 配送员姓名
    PostmanName   string `json:"postman_name,omitempty"`

    // 配送员编码
    PostmanCode   string `json:"postman_code,omitempty"`

    // 配送员手机
    PostmanPhone   string `json:"postman_phone,omitempty"`

    // 服务商名称
    ProviderName   string `json:"provider_name,omitempty"`

    // 服务商编码
    ProviderCode   string `json:"provider_code,omitempty"`

}
