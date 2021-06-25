package user

// OpenAccountTokenValidateResult 
type OpenAccountTokenValidateResult struct {

    // token中的数据
    Data   *TokenInfo `json:"data,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

    // 是否成功
    Successful   bool `json:"successful,omitempty"`

    // 错误码
    Code   int64 `json:"code,omitempty"`

}