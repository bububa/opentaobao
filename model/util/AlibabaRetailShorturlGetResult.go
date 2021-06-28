package util

// AlibabaRetailShorturlGetResult 
/* model for simplify = false
type AlibabaRetailShorturlGetResult struct {

    // module
    
    Module  *struct {
        ShortUrlDto  *ShortUrlDto `json:"short_url_dto,omitempty"`
    } `json:"module,omitempty"`
    

    // errorInfos
    
    ErrorInfos  struct {
        ErrorInfo  []ErrorInfo `json:"error_info,omitempty"`
    } `json:"error_infos,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaRetailShorturlGetResult 
type AlibabaRetailShorturlGetResult struct {

    // module
    Module   *ShortUrlDto `json:"module,omitempty"`

    // errorInfos
    ErrorInfos   []ErrorInfo `json:"error_infos,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

}
