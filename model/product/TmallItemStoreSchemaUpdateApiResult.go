package product

// TmallItemStoreSchemaUpdateApiResult 
type TmallItemStoreSchemaUpdateApiResult struct {

    // 错误信息
    ErMsg   string `json:"er_msg,omitempty"`

    // 错误码
    ErCode   string `json:"er_code,omitempty"`

    // 商品ID
    Result   int64 `json:"result,omitempty"`

    // 成功
    Error   bool `json:"error,omitempty"`

    // 错误信息
    MappedErrorMessages   string `json:"mapped_error_messages,omitempty"`

}
